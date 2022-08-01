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

}

func generateRequest(request *Request) (s *Statement, err error) {
	category := request.Category
	name := request.RequestType

	fmt.Printf("Request %s %s\n", category, name)

	var structName string
	s = Line()

	// Params
	structName = name + "Params"
	s.Commentf("%s represents the params body for the %q request.\n%s\n", structName, name, request.Description).Line()

	if err = generateStructFromParams("request", s, structName, request.RequestFields); err != nil {
		return nil, fmt.Errorf("Failed parsing 'Params' for request %q in category %q", name, category)
	}

	// Returns
	structName = name + "Response"
	s.Commentf("%s represents the response body for the %q request.\n%s\n", structName, name, request.Description).Line()

	if err = generateStructFromParams("response", s, structName, request.ResponseFields); err != nil {
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

	s.Commentf("%s sends the corresponding request to the connected OBS WebSockets server.", name).Do(func(z *Statement) {
		if hasRequiredArgs {
			return
		}
		z.Id("Note the variadic arguments as this request doesn't require any parameters.")
	})
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
		List(Id("resp"), Id("err")).Op(":=").Id("c").Dot("SendRequest").Call(
			Lit(name),
			Id("params"),
		),
		If(Id("err").Op("!=").Nil()).Block(
			Return().List(Nil(), Id("err")),
		),
		Return().List(Id("resp").Assert(Op("*").Id(name+"Response")), Nil()),
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
	f := NewFile("goobs")
	f.HeaderComment("This file has been automatically generated. Don't edit it.")
	f.Add(
		Func().Id("GetEventForType").Params(Id("name").String()).Interface().Block(
			Switch(Id("name")).BlockFunc(func(g *Group) {
				for _, e := range events {
					g.Case(Lit(e.EventType))
					g.Return(Op("&").Qual(goobs+"/api/events", e.EventType).Values())
				}
				g.Default().Return(Nil())
			}),
		),
	)
	if err := f.Save(fmt.Sprintf("%s/zz_generated.events.go", root)); err != nil {
		panic(err)
	}
}

func generateEvent(event *Event) (s *Statement, err error) {
	fmt.Printf("Event %s %s\n", event.Category, event.EventType)

	s = Line()
	s.Commentf("%s represents the event body for the %q event.\nSince v%s.", event.EventType, event.EventType, event.InitialVersion).Line()

	if err = generateStructFromParams("event", s, event.EventType, event.DataFields); err != nil {
		return nil, fmt.Errorf("Failed generating event %q in category %q", event.EventType, event.Category)
	}

	return s, nil
}

func generateEventSubscriptions(enums []*Enum) {
	dir := fmt.Sprintf("%s/api/events/subscriptions", root)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}

	for _, e := range enums {
		if e.EnumType != "EventSubscription" {
			continue
		}

		f := NewFile("subscriptions")
		f.HeaderComment("This file has been automatically generated. Don't edit it.")

		for _, subscription := range e.EnumIdentifiers {
			val := ""
			switch value := subscription.EnumValue.(type) {
			case string:
				val = value
			default:
				val = fmt.Sprintf("%v", value)
			}

			id := subscription.EnumIdentifier
			description := subscription.Description

			fmt.Printf("EventSubscription %s\n", id)

			f.Add(
				Comment(description),
				Line(),
				Const().Id(id).Op("=").Id(val),
			)
		}

		if err := f.Save(fmt.Sprintf("%s/xx_generated.subscriptions.go", dir)); err != nil {
			panic(err)
		}
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

		//fieldName, err := sanitizeText(field.Name)
		//if err != nil {
		//	return fmt.Errorf("Failed sanitizing field %#v", field)
		//}

		var fieldType *Statement
		switch fvt {
		case "String":
			fieldType = String()
		case "Array<String>":
			fieldType = Index().String()
		case "Number":
			fieldType = Float64()
		case "Array<Object>":
			fieldType = Index().Interface()
		case "Object":
			fieldType = Interface()
		case "Boolean":
			fieldType = Bool()
		case "Any":
			fieldType = Interface()
		default:
			panic(fmt.Errorf("in struct %q, %q is of weird type %q", name, fvn, fvt))
		}

		keysInfo[fvn] = keyInfo{
			Type:      fieldType,
			Comment:   fvd,
			NoJSONTag: false,
			Embedded:  embedded,
			OmitEmpty: true,
		}
	}

	//fmt.Printf("%#v\n", keysInfo)
	statement, err := parseJenKeysAsStruct(name, keysInfo)
	if err != nil {
		return fmt.Errorf("Failed parsing dotted key: %s", err)
	}

	s.Add(statement)

	return nil
}
