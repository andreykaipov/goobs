package main

import (
	"fmt"
	"os"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func snake(s string) string {
	return strings.ReplaceAll(s, " ", "_")
}
func lower(s string) string {
	return strings.ToLower(s)
}

func generateRequests(requests []*Request) {
	topClientFields := []Code{}
	topClientSetters := []Code{}

	// aggregate requests by category
	categories := map[string][]*Request{}
	for _, request := range requests {
		categories[request.Category] = append(categories[request.Category], request)
	}

	for _, category := range sortedKeys(categories) {
		requestsInCategory := categories[category]

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
		client.Commentf("Client represents a client for '%s' requests.", category)
		client.Add(
			Type().Id("Client").Struct(
				Op("*").Qual(goobs+"/api", "Client"),
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
		for _, request := range requestsInCategory {
			name := request.RequestType

			if request.Deprecated {
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
			fName := strings.ToLower(name)
			if err := f.Save(fmt.Sprintf("%s/xx_generated.%s.go", dir, fName)); err != nil {
				panic(err)
			}
		}
	}

	var f *File

	// Write utils for the top-level client
	f = NewFile("goobs")
	f.HeaderComment("This file has been automatically generated. Don't edit it.")
	f.Add(Type().Id("subclients").Struct(topClientFields...))
	f.Add(Func().Id("setClients").Params(Id("c").Op("*").Id("Client")).Block(topClientSetters...))
	if err := f.Save(fmt.Sprintf("%s/zz_generated.client.go", root)); err != nil {
		panic(err)
	}

	/*
		// GetRequestResponseForType to get the Go type of responses from their names
		f = NewFile("goobs")
		f.HeaderComment("This file has been automatically generated. Don't edit it.")
		f.Add(
			Func().Id("GetRequestResponseForType").Params(Id("name").String()).Interface().Block(
				Switch(Id("name")).BlockFunc(func(g *Group) {
					for _, category := range sortedKeys(categories) {
						categorySnake := strings.ReplaceAll(category, " ", "_")
						for _, request := range categories[category] {
							g.Case(Lit(request.RequestType))
							g.Return(Op("&").Qual(goobs+"/api/requests/"+categorySnake, request.RequestType+"Response").Values())
						}
					}
					g.Default().Return(Nil())
				}),
			),
		)
		if err := f.Save(fmt.Sprintf("%s/zz_generated.requests.go", root)); err != nil {
			panic(err)
		}
	*/
}

func generateRequest(request *Request) (s *Statement, err error) {
	category := request.Category
	name := request.RequestType

	fmt.Printf("Request %s %s\n", category, name)

	var structName string
	s = Line()

	// Params
	structName = name + "Params"
	s.Commentf("Represents the request body for the %s request.", name).Line()

	if err := generateStructFromParams("request", s, structName, request.RequestFields); err != nil {
		return nil, fmt.Errorf("Failed parsing 'Params' for request %q in category %q", name, category)
	}

	s.Add(
		Commentf("Returns the associated request.").Line(),
		Func().Params(Id("o").Op("*").Id(structName)).Id("GetRequestName").Params().String().Block(
			Return(Lit(request.RequestType)),
		).Line(),
	)

	// Returns
	structName = name + "Response"
	s.Commentf("Represents the response body for the %s request.", name).Line()

	if err := generateStructFromParams("response", s, structName, request.ResponseFields); err != nil {
		return nil, fmt.Errorf("Failed parsing 'Returns' for request %q in category %q", name, category)
	}

	// generate the request function

	allOptional := true
	for _, f := range request.RequestFields {
		allOptional = allOptional && f.ValueOptional
	}

	// we should use variadic args if we have no RequestFields in our
	// params, or if all of the other fields above were optional
	varargs := len(request.RequestFields) == 0 || allOptional
	hasRequiredArgs := !varargs

	s.Commentf("%s", request.Description)
	s.Line()
	s.Func().Params(Id("c").Op("*").Id("Client")).Id(name).ParamsFunc(func(g *Group) {
		if hasRequiredArgs {
			g.Id("params").Op("*").Id(name + "Params")
		} else {
			g.Id("paramss").Op("...").Op("*").Id(name + "Params")
		}
	}).Params(Op("*").Id(name+"Response"), Error()).Block(
		Do(func(z *Statement) {
			if hasRequiredArgs {
				return
			}
			z.If(Len(Id("paramss")).Op("==").Lit(0)).Block(
				Id("paramss").Op("=").Index().Op("*").Id(name + "Params").Values(Values()),
			)
			z.Line()
			z.Id("params").Op(":=").Id("paramss").Index(Lit(0))
		}),
		Id("data").Op(":=").Op("&").Id(name+"Response").Values(),
		Return().List(Id("data"), Id("c").Dot("SendRequest").Call(Id("params"), Id("data"))),
	)

	return s, nil
}

func generateEvents(events []*Event) {
	// Unlike requests, I don't want to write each event into a package
	// representing the category, since I want the syntax to read like
	// `*events.TransitionBegin` instead of `*transitions.TransitionBegin`,
	// when reading from the eventing loop.
	dir := fmt.Sprintf("%s/api/events", root)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}

	for _, e := range events {
		category := snake(e.Category)
		name := e.EventType

		s, err := generateEvent(e)
		if err != nil {
			panic(err)
		}

		f := NewFile("events")
		f.HeaderComment("This file has been automatically generated. Don't edit it.")
		f.Add(s)
		fName := lower(name)
		if err := f.Save(fmt.Sprintf("%s/xx_generated.%s.%s.go", dir, category, fName)); err != nil {
			panic(err)
		}
	}

	// GetEventForType to get the Go type of responses from their names
	f := NewFile("events")
	f.HeaderComment("This file has been automatically generated. Don't edit it.")
	f.Add(
		Func().Id("GetType").Params(Id("name").String()).Interface().Block(
			Switch(Id("name")).BlockFunc(func(g *Group) {
				for _, e := range events {
					g.Case(Lit(e.EventType))
					g.Return(Op("&").Id(e.EventType).Values())
				}
				g.Default().Return(Nil())
			}),
		),
	)
	if err := f.Save(fmt.Sprintf("%s/api/events/zz_generated.events.go", root)); err != nil {
		panic(err)
	}
}

func generateEvent(event *Event) (s *Statement, err error) {
	fmt.Printf("Event %s %s\n", event.Category, event.EventType)

	s = Line()
	s.Commentf("Represents the event body for the %s event.", event.EventType).Line()

	if err := generateStructFromParams("event", s, event.EventType, event.DataFields); err != nil {
		return nil, fmt.Errorf("Failed generating event %q in category %q", event.EventType, event.Category)
	}

	return s, nil
}

type enumFilter func(e *Enum) bool

func generateEventSubscriptions(enums []*Enum, filter enumFilter) {
	dir := fmt.Sprintf("%s/api/events/subscriptions", root)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}

	s := Line()

	for _, e := range enums {
		if !filter(e) {
			continue
		}

		f := NewFile("subscriptions")
		f.HeaderComment("This file has been automatically generated. Don't edit it.")
		f.Add(s)

		s.Const().DefsFunc(func(g *Group) {
			for _, x := range e.EnumIdentifiers {
				id := x.EnumIdentifier
				fmt.Printf("EventSubscription %s\n", id)

				val := ""
				switch value := x.EnumValue.(type) {
				case string:
					val = value
				default:
					val = fmt.Sprintf("%v", value)
				}

				g.Comment(x.Description)
				g.Id(id).Op("=").Id(val)
				g.Line()
			}
		})

		if err := f.Save(fmt.Sprintf("%s/xx_generated.subscriptions.go", dir)); err != nil {
			panic(err)
		}

		return
	}
}

func generateRequestStatuses(enums []*Enum, filter enumFilter) {
	dir := fmt.Sprintf("%s/api/requests", root)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}

	s := Line()

	for _, e := range enums {
		if !filter(e) {
			continue
		}

		f := NewFile("requests")
		f.HeaderComment("This file has been automatically generated. Don't edit it.")

		s.Func().Id("GetStatusForCode").Params(Id("code").Int()).String().Block(
			Switch(Id("code")).BlockFunc(func(g *Group) {
				for _, x := range e.EnumIdentifiers {
					id := x.EnumIdentifier
					fmt.Printf("RequestStatus %s\n", id)

					val := -69
					switch value := x.EnumValue.(type) {
					// json decoder will parse all numbers as floats
					case float64:
						val = int(value)
					default:
						panic(fmt.Errorf("enum %s has value %#[2]v not an int %[2]T", id, value))
					}

					g.Case(Lit(val))
					g.Comment(x.Description)
					g.Return(Lit(id))
				}
				g.Default().Return(Lit("EvenMoreUnknown"))
			}),
		)

		f.Add(s)

		if err := f.Save(fmt.Sprintf("%s/xx_generated.request_status.go", dir)); err != nil {
			panic(err)
		}

		return
	}
}

func generateStructFromParams(origin string, s *Statement, name string, fields interface{}) error {
	var fs []Field

	switch v := fields.(type) {
	case []*DataField:
		fs = make([]Field, len(v))
		for i := range fs {
			fs[i] = v[i]
		}
	case []*RequestField:
		fs = make([]Field, len(v))
		for i := range fs {
			fs[i] = v[i]
		}
	case []*ResponseField:
		fs = make([]Field, len(v))
		for i := range fs {
			fs[i] = v[i]
		}
	default:
		panic("uh")
	}

	keysInfo := map[string]keyInfo{}

	for _, f := range fs {
		fvn := f.GetValueName()
		fvt := f.GetValueType()
		fvd := f.GetValueDescription()
		embedded := false

		var fieldType *Statement
		switch fvt {
		case "String":
			fieldType = String()
		case "Array<String>":
			fieldType = Index().String()
		case "Number":
			fieldType = Float64()
		case "Boolean":
			switch origin {
			case "request":
				fieldType = Op("*").Bool()
			default:
				fieldType = Bool()
			}
		case "Any":
			fieldType = Interface()
		case "Object":
			switch fvn {
			case "inputSettings":
				fallthrough
			case "defaultInputSettings":
				fallthrough
			case "filterSettings":
				fallthrough
			case "defaultFilterSettings":
				fallthrough
			case "transitionSettings":
				fieldType = Map(String()).Interface()
			default:
				fieldType = Interface()
			}
		case "Array<Object>":
			switch fvn {
			default:
				fieldType = Index().Interface()
			}
		default:
			panic(fmt.Errorf("in struct %q, %q is of weird type %q", name, fvn, fvt))
		}

		if key, keyInfo := handleCommonObjects(origin, fvn, name); keyInfo != nil {
			fmt.Printf("  > %-25s handled as (or part of) common struct %s\n", fvn, key)
			keysInfo[key] = *keyInfo
			continue
		}

		keysInfo[fvn] = keyInfo{
			Type:      fieldType,
			Comment:   fvd,
			NoJSONTag: false,
			Embedded:  embedded,
			OmitEmpty: true,
		}
	}

	statement, err := parseJenKeysAsStruct(name, keysInfo)
	if err != nil {
		return fmt.Errorf("Failed parsing dotted key: %s", err)
	}

	s.Add(statement)

	return nil
}

func handleCommonObjects(origin, fieldName, structName string) (string, *keyInfo) {
	type nestedInfo struct {
		Refer   string
		Comment string
		OnlyFor []string // additional requirement to allow us to use different types for the same key prefix
	}

	// key prefix to manually declared struct in typedefs package
	lookup := map[string]nestedInfo{
		"keyModifiers.":         {"KeyModifiers", "Key modifiers to apply", nil},
		"streamServiceSettings": {"StreamServiceSettings", " ", nil},
		"inputAudioTracks":      {"InputAudioTracks", "", nil},
		"sceneItemTransform":    {"SceneItemTransform", "Scene item transform info", nil},
		"monitors":              {"[]Monitor", "List of detected monitors", nil},
		"scenes":                {"[]Scene", "", nil},
		"sceneItems":            {"[]SceneItem", "", nil},
		"inputs":                {"[]Input", "", nil},
		"filters":               {"[]Filter", "", nil},
		"transitions":           {"[]Transition", "", nil},
		"propertyItems":         {"[]PropertyItem", "", nil},
		// "settings.":             {"StreamSettings", " ", []string{"GetStreamSettings", "SetStreamSettings"}},
		// "stream.settings.":      {"StreamSettings", " ", []string{}},
	}

	valid := func(i nestedInfo) bool {
		if len(i.OnlyFor) == 0 {
			return true
		}

		for _, v := range i.OnlyFor {
			if strings.HasPrefix(structName, v) {
				return true
			}
		}

		return false
	}

	for k, info := range lookup {
		if !valid(info) {
			continue
		}

		if strings.HasPrefix(fieldName, k) {
			k = strings.TrimSuffix(k, ".")

			s := Null()
			if strings.HasPrefix(info.Refer, "[]") {
				s = Index()
				info.Refer = strings.TrimPrefix(info.Refer, "[]")
			}

			s = s.Op("*").Qual(typedefQualifier(origin), info.Refer)

			return k, &keyInfo{Type: s, Comment: info.Comment, OmitEmpty: true}
		}
	}

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
