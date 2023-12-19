package main

import (
	"fmt"
	"os"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func snake(s string) string { return strings.ReplaceAll(s, " ", "_") }
func trim(s string) string  { return strings.ReplaceAll(s, " ", "") }
func lower(s string) string { return strings.ToLower(s) }

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

		categoryPascal := pascal(category)
		categoryClaustrophic := trim(category)

		// For the top-level client
		qualifier := goobs + "/api/requests/" + categoryClaustrophic
		topClientFields = append(topClientFields, Id(categoryPascal).Op("*").Qual(qualifier, "Client"))
		topClientSetters = append(
			topClientSetters, Id("c").Dot(categoryPascal).Op("=").Op("&").Qual(qualifier, "Client").Values(
				Id("Client").Op(":").Id("c").Dot("Client"),
			),
		)

		// Generate the category-level client
		client := NewFile(categoryClaustrophic)
		client.HeaderComment("This file has been automatically generated. Don't edit it.")
		client.Add(
			Type().Id("_response").Op("=").Qual(goobs+"/api", "ResponseCommon"),
			Line(),
			Commentf("Client represents a client for '%s' requests.", category),
			Line(),
			Type().Id("Client").Struct(
				Op("*").Qual(goobs+"/api", "Client"),
			),
		)

		// Write the category-level client
		dir := fmt.Sprintf("%s/api/requests/%s", root, categoryClaustrophic)
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
			fName := lower(name)
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

	respf := &ResponseField{}
	respf.ValueName = "_response"
	respf.ValueType = "~requests~" // internal type
	request.ResponseFields = append(request.ResponseFields, respf)

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
		category := trim(e.Category)
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
		Func().Id("GetType").Params(Id("name").String()).Any().Block(
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
	s.Commentf("Represents the event body for the %s event.\n\n%s", event.EventType, event.Description).Line()

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

func generateStructFromParams[F Field](origin string, s *Statement, name string, fields []F) error {
	keysInfo := map[string]keyInfo{}

	for _, f := range fields {
		fvn := f.GetValueName()
		fvt := f.GetValueType()
		fvd := f.GetValueDescription()

		// some fields document the type of the fields in the Object
		// like keyModifiers.shift, but we handle these in our manually
		// written typedefs
		if strings.Contains(fvn, ".") {
			fmt.Printf("generateStructFromParams for %s/%s: skipping %s\n", origin, name, fvn)
			continue
		}

		embedded := false

		var fieldType *Statement
		switch fvt {
		case "String":
			switch origin {
			case "request":
				fieldType = Op("*").String()
			default:
				fieldType = String()
			}
		case "Number":
			fieldType = Float64()
			if strings.HasSuffix(fvn, "Id") || strings.HasSuffix(fvn, "Index") {
				fieldType = Int()
			}
			switch origin {
			case "request":
				fieldType = Op("*").Add(fieldType)
			}
		case "Boolean":
			switch origin {
			case "request":
				fieldType = Op("*").Bool()
			default:
				fieldType = Bool()
			}
		case "Array<String>":
			fieldType = Index().String()
		case "Array<Number>":
			fieldType = Index().Float64()
		case "Array<Boolean>":
			fieldType = Index().Bool()
		case "Any":
			fieldType = Any()
		case "Object":
			fieldType = mapObject(origin, name, f)
		case "Array<Object>":
			fieldType = mapArrayObject(origin, name, f)
		case "~requests~":
			fieldType = Id(fvn)
			embedded = true
		default:
			panic(fmt.Errorf("in struct %q, %q is of weird type %q", name, fvn, fvt))
		}

		keysInfo[fvn] = keyInfo{
			Type:          fieldType,
			Comment:       fvd,
			NoJSONTag:     false,
			Embedded:      embedded,
			OmitEmpty:     true,
			ExposeBuilder: origin == "request",
		}
	}

	statement, err := parseJenKeysAsStruct(name, keysInfo)
	if err != nil {
		return fmt.Errorf("Failed parsing dotted key: %s", err)
	}

	s.Add(statement)

	return nil
}
