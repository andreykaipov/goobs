package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"unicode"

	. "github.com/dave/jennifer/jen"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func generateRequests(data *Comments) {
	topClientFields := []Code{}
	topClientSetters := []Code{}

	for _, category := range sortedKeys(data.Requests) {
		fmt.Printf("- %s \n", category)

		categorySnake := strings.ReplaceAll(category, " ", "_")
		categoryPascal := strings.ReplaceAll(strings.Title(category), " ", "")
		categoryClaustrophic := strings.ReplaceAll(category, " ", "")

		// For the top-level client
		qualifier := goobs + "/api/requests/" + categorySnake
		topClientFields = append(topClientFields, Id(categoryPascal).Op("*").Qual(qualifier, "Client"))
		topClientSetters = append(
			topClientSetters, Id("c").Dot(categoryPascal).Op("=").Op("&").Qual(qualifier, "Client").Values(
				Id("Client").Op(":").Id("c").Dot("Client"),
			),
		)

		// Generate the category-level client
		client := NewFile(categoryClaustrophic)
		client.HeaderComment("This file has been automatically generated. Don't edit it.")
		client.Commentf("Client represents a client for '%s' requests", category)
		client.Add(
			Type().Id("Client").Struct(
				Op("*").Qual(goobs+"/api/requests", "Client"),
			),
		)

		// Write the category-level client
		dir := fmt.Sprintf("%s/api/requests/%s", root, categorySnake)
		if err := os.MkdirAll(dir, 0777); err != nil {
			panic(err)
		}
		if err := client.Save(fmt.Sprintf("%s/zz_generated.client.go", dir)); err != nil {
			panic(err)
		}

		// Generate the requests for the category
		for _, request := range data.Requests[category] {
			// fmt.Printf("  - %s\n", request.Name)

			if request.Deprecated != "" {
				// fmt.Fprintf(os.Stderr, "Request %q is deprecated\n", request.Name)
				continue
			}

			s, err := generateRequest(request)
			if err != nil {
				panic(err)
			}

			f := NewFile(categoryClaustrophic)
			f.HeaderComment("This file has been automatically generated. Don't edit it.")
			f.Add(s)
			fName := strings.ToLower(request.Name)
			if err := f.Save(fmt.Sprintf("%s/xx_generated.%s.go", dir, fName)); err != nil {
				panic(err)
			}
		}
	}

	// Write utils for the top-level client
	f := NewFile("goobs")
	f.HeaderComment("This file has been automatically generated. Don't edit it.")
	f.Add(Type().Id("subclients").Struct(topClientFields...))
	f.Add(Func().Id("setClients").Params(Id("c").Op("*").Id("Client")).Block(topClientSetters...))

	if err := f.Save(fmt.Sprintf("%s/zz_generated.client.go", root)); err != nil {
		panic(err)
	}
}

func generateRequest(request *Request) (s *Statement, err error) {
	var structName string
	s = Line()

	// Params
	structName = request.Name + "Params"
	s.Commentf("%s represents the params body for the %q request.\n%s%s\nSince %s.", structName, request.Name, request.Lead, request.Description, request.Since).Line()
	request.Params = append(request.Params, &Param{Name: "ParamsBasic", Type: "~requests~"}) // internal type
	if err = generateStructFromParams("request", s, structName, request.Params); err != nil {
		return nil, fmt.Errorf("Failed parsing 'Params' for request %q in category %q", request.Name, request.Category)
	}

	s.Add(
		Commentf("GetSelfName just returns %q.", request.Name).Line(),
		Func().Params(Id("o").Op("*").Id(structName)).Id("GetSelfName").Params().String().Block(
			Return(Lit(request.Name)),
		).Line(),
	)

	// Returns
	structName = request.Name + "Response"
	s.Commentf("%s represents the response body for the %q request.\n%s%s\nSince v%s.", structName, request.Name, request.Lead, request.Description, request.Since).Line()
	request.Returns = append(request.Returns, &Param{Name: "ResponseBasic", Type: "~requests~"}) // internal type
	if err = generateStructFromParams("response", s, structName, request.Returns); err != nil {
		return nil, fmt.Errorf("Failed parsing 'Returns' for request %q in category %q", request.Name, request.Category)
	}

	// generate the request function

	hasRequiredArgs := len(request.Params) > 1

	s.Commentf("%s sends the corresponding request to the connected OBS WebSockets server.", request.Name).Do(func(z *Statement) {
		if hasRequiredArgs {
			return
		}
		z.Id("Note the variadic arguments as this request doesn't require any parameters.")
	})
	s.Line()
	s.Func().Params(Id("c").Op("*").Id("Client")).Id(request.Name).ParamsFunc(func(g *Group) {
		if hasRequiredArgs {
			g.Id("params").Op("*").Id(request.Name + "Params")
		} else {
			g.Id("paramss").Op("...").Op("*").Id(request.Name + "Params")
		}
	}).Params(Op("*").Id(request.Name+"Response"), Error()).Block(
		Do(func(z *Statement) {
			if hasRequiredArgs {
				return
			}
			z.If(Len(Id("paramss")).Op("==").Lit(0)).Block(
				Id("paramss").Op("=").Index().Op("*").Id(request.Name + "Params").Values(Values()),
			)
			z.Line()
			z.Id("params").Op(":=").Id("paramss").Index(Lit(0))
		}),
		Id("data").Op(":=").Op("&").Id(request.Name+"Response").Values(),
		If(
			Id("err").Op(":=").Id("c").Dot("SendRequest").Call(
				Id("params"), Id("data"),
			),
			Id("err").Op("!=").Nil(),
		).Block(
			Return().List(Nil(), Id("err")),
		),
		Return().List(Id("data"), Nil()),
	)

	return s, nil
}

func generateEvents(data *Comments) {
	events := []*Event{}

	// Unlike requests, I don't want to write each event into a package
	// representing the category, since I want the syntax to read like
	// `*events.TransitionBegin` instead of `*transitions.TransitionBegin`,
	// when reading from the eventing loop.
	dir := fmt.Sprintf("%s/api/events", root)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}

	for _, category := range sortedKeys(data.Events) {
		fmt.Printf("- %s\n", category)

		categorySnake := strings.ReplaceAll(category, " ", "_")

		// Generate the events for the category
		for _, event := range data.Events[category] {
			// fmt.Printf("  - %s\n", event.Name)

			s, err := generateEvent(event)
			if err != nil {
				panic(err)
			}

			f := NewFile("events")
			f.HeaderComment("This file has been automatically generated. Don't edit it.")
			f.Add(s)
			fName := strings.ToLower(event.Name)
			if err := f.Save(fmt.Sprintf("%s/xx_generated.%s.%s.go", dir, categorySnake, fName)); err != nil {
				panic(err)
			}

			events = append(events, event)
		}
	}

	f := NewFile("events")
	f.HeaderComment("This file has been automatically generated. Don't edit it.")
	f.Add(
		Func().Id("GetEventForType").Params(Id("name").String()).Id("Event").Block(
			Switch(Id("name")).BlockFunc(func(g *Group) {
				for _, e := range events {
					g.Case(Lit(e.Name))
					g.Return(Op("&").Id(e.Name).Values())
				}
				g.Default().Return(Nil())
			}),
		),
	)
	if err := f.Save(fmt.Sprintf("%s/zz_generated.events.go", dir)); err != nil {
		panic(err)
	}
}

func generateEvent(event *Event) (s *Statement, err error) {
	s = Line()
	s.Commentf("%s represents the event body for the %q event.\nSince v%s.", event.Name, event.Name, event.Since).Line()

	event.Returns = append(event.Returns, &Param{Name: "EventBasic", Type: "~events~"}) // internal type
	if err = generateStructFromParams("event", s, event.Name, event.Returns); err != nil {
		return nil, fmt.Errorf("Failed generating event %q in category %q", event.Name, event.Category)
	}

	return s, nil
}

func generateTypeDefs(data *Comments) {
	dir := fmt.Sprintf("%s/api/typedefs", root)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}

	for _, typeDef := range data.TypeDefs {
		name := typeDef.TypeDefs[0].Name
		fmt.Printf("- %s\n", name)

		s, err := generateTypeDef(typeDef)
		if err != nil {
			panic(err)
		}

		f := NewFile("typedefs")
		f.HeaderComment("This file has been automatically generated. Don't edit it.")
		f.Add(s)
		fName := strings.ToLower(name)
		if err := f.Save(fmt.Sprintf("%s/xx_generated.%s.go", dir, fName)); err != nil {
			panic(err)
		}
	}
}

func generateTypeDef(typeDef *TypeDef) (s *Statement, err error) {
	name := typeDef.TypeDefs[0].Name

	s = Line()
	s.Commentf("%s represents the complex type for %s.", name, name).Line()

	if err = generateStructFromParams("typedef", s, name, typeDef.Properties); err != nil {
		return nil, fmt.Errorf("Failed generating struct for complex type %q", name)
	}

	return s, nil
}

func generateStructFromParams(origin string, s *Statement, name string, params []*Param) error {
	keysInfo := map[string]keyInfo{}
	noOptional := regexp.MustCompile(`(?i)[ ]+?\(optional\)([ ]+)?$`)

	for _, field := range params {
		fieldName, err := sanitizeText(field.Name)
		if err != nil {
			return fmt.Errorf("Failed sanitizing field %#v", field)
		}

		noJSONTag := false
		embedded := false

		var fieldType *Statement
		switch val := noOptional.ReplaceAllString(field.Type, ""); val {
		case "string":
			fieldType = String()
		case "String":
			fieldType = String()
		case "String | Object":
			switch fieldName {
			case "item": // this is ok, it's implicitly an object by the keys following it
			default:
				panic(fmt.Errorf("in struct %q, %q can't be both %q", name, field.Name, field.Type))
			}
		case "int":
			fieldType = Int()
		case "Integer":
			fieldType = Int()
		case "double":
			fieldType = Float64()
		case "float":
			fieldType = Float64()
		case "Number":
			// only happens in SceneItem https://github.com/Palakis/obs-websocket/blob/4.9.1/src/Utils.cpp#L160
			// used for both Ints & Floats, so we'll use Float
			fieldType = Float64()
		case "bool":
			fieldType = Bool()
		case "boolean":
			fieldType = Bool()
		case "Boolean":
			fieldType = Bool()
		// funkier types
		case "Object":
			fieldType = Map(String()).Interface()
		case "Array<String>":
			fieldType = Index().String()
		case "Array<Object>":
			fieldType = Index().Map(String()).Interface()
		// all the possible typedefs
		case "SceneItem":
			fallthrough
		case "SceneItemTransform":
			fallthrough
		case "OBSStats":
			fallthrough
		case "Output":
			fallthrough
		case "ScenesCollection":
			fallthrough
		case "Scene":
			fieldType = Qual(typedefQualifier(origin), val)
		// arrays of the all the possible typedefs
		case "Array<SceneItem>":
			fallthrough
		case "Array<SceneItemTransform>":
			fallthrough
		case "Array<OBSStats>":
			fallthrough
		case "Array<Output>":
			fallthrough
		case "Array<ScenesCollection>":
			fallthrough
		case "Array<Scene>":
			typedef := val[6 : len(val)-1]
			fieldType = Index().Qual(typedefQualifier(origin), typedef)
		// internal types we use in our generation to embed base structs
		case "~requests~":
			fieldType = Qual(goobs+"/api/requests", field.Name)
			embedded = true
		case "~events~":
			fieldType = Id(field.Name)
			embedded = true
		default:
			panic(fmt.Errorf("in struct %q, %q is of weird type %q", name, field.Name, field.Type))
		}

		if key, keyInfo := handleCommonObjects(origin, fieldName); keyInfo != nil {
			keysInfo[key] = *keyInfo
			continue
		}

		keysInfo[fieldName] = keyInfo{
			Type:      fieldType,
			Comment:   field.Description,
			NoJSONTag: noJSONTag,
			Embedded:  embedded,
		}
	}

	statement, err := parseJenKeysAsStruct(name, keysInfo)
	if err != nil {
		return fmt.Errorf("Failed parsing dotted key: %s", err)
	}

	s.Add(statement)

	return nil
}

func handleCommonObjects(origin, fieldName string) (string, *keyInfo) {
	if !strings.Contains(fieldName, ".") {
		return "", nil
	}

	// key prefix to manually declared struct in typedefs package
	lookup := map[string][]string{
		"bounds.":          {"Bounds", "The bounding box of the object (source, scene item, etc)."},
		"crop.":            {"Crop", "The crop specification for the object (source, scene item, etc)."},
		"position.":        {"Position", "The position of the object (source, scene item, etc)."},
		"scale.":           {"Scale", "The scaling specification for the object (source, scene item, etc)."},
		"item.":            {"Item", "The item specification for this object."},
		"font.":            {"Font", "The font specification for this object."},
		"settings.":        {"Settings", " "},
		"stream.settings.": {"Settings", " "},
		"items.*.":         {"Item", "The items for this object."},
		"filters.*.":       {"Filter", "The filters for this object."},
	}

	for k, vs := range lookup {
		if strings.HasPrefix(fieldName, k) {
			k = strings.TrimSuffix(k, ".")
			s := Qual(typedefQualifier(origin), vs[0])
			return k, &keyInfo{Type: s, Comment: vs[1]}
		}
	}

	fmt.Printf("  > not handled as a common struct: %s\n", fieldName)

	return "", nil
}

// If the origin of the generation is from a typedef, we don't actually need
// a Qualifier. Saves us repeating this switch/case a bunch
func typedefQualifier(origin string) string {
	switch origin {
	case "typedef":
		return ""
	default:
		return goobs + "/api/typedefs"
	}
}

func sanitizeText(text string) (string, error) {
	isMn := func(r rune) bool {
		// Mn: nonspacing marks
		return unicode.Is(unicode.Mn, r)
	}

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	clean, _, err := transform.String(t, text)
	return clean, err
}

func sortedKeys(m interface{}) []string {
	keys := reflect.ValueOf(m).MapKeys()

	sorted := make([]string, len(keys))
	i := 0
	for _, key := range keys {
		sorted[i] = key.Interface().(string)
		i++
	}
	sort.Strings(sorted)

	return sorted
}
