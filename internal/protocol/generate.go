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

func generateEvents(events []*Event) {
	// Unlike requests, I don't want to write each event into a package
	// representing the category, since I want the syntax to read like
	// `*events.TransitionBegin` instead of `*transitions.TransitionBegin`,
	// when reading from the eventing loop.
	dir := fmt.Sprintf("%s/apiv5/events", root)
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}

	for _, e := range events {
		category := snake(e.Category)

		fmt.Printf("%-12s %s\n", e.Category, e.EventType)

		s, err := generateEvent(e)
		if err != nil {
			panic(err)
		}

		f := NewFile("events")
		f.HeaderComment("This file has been automatically generated. Don't edit it.")
		f.Add(s)
		fName := lower(e.EventType)
		if err := f.Save(fmt.Sprintf("%s/xx_generated.%s.%s.go", dir, category, fName)); err != nil {
			panic(err)
		}
	}

	f := NewFile("events")
	f.HeaderComment("This file has been automatically generated. Don't edit it.")
	f.Add(
		Func().Id("GetEventForType").Params(Id("name").String()).Id("Event").Block(
			Switch(Id("name")).BlockFunc(func(g *Group) {
				for _, e := range events {
					g.Case(Lit(e.EventType))
					g.Return(Op("&").Id(e.EventType).Values())
				}
				g.Default().Return(Nil())
			}),
		),
	)
	if err := f.Save(fmt.Sprintf("%s/zz_generated.events.go", dir)); err != nil {
		panic(err)
	}
}

func generateEventSubscriptions(enums []*Enum) {
	dir := fmt.Sprintf("%s/apiv5/events/subscriptions", root)
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

			fmt.Printf("%-12s %s\n", "Subscription", id)

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

func generateEvent(event *Event) (s *Statement, err error) {
	s = Line()
	s.Commentf("%s represents the event body for the %q event.\nSince v%s.", event.EventType, event.EventType, event.InitialVersion).Line()

	//event.Returns = append(event.Returns, &Param{Name: "EventBasic", Type: "~events~"}) // internal type
	if err = generateStructFromParams("event", s, event.EventType, event.DataFields); err != nil {
		return nil, fmt.Errorf("Failed generating event %q in category %q", event.EventType, event.Category)
	}

	return s, nil
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
		default:
			panic(fmt.Errorf("in struct %q, %q is of weird type %q", name, fvn, fvt))
		}

		keysInfo[fvn] = keyInfo{
			Type:      fieldType,
			Comment:   fvd,
			NoJSONTag: false,
			Embedded:  false,
			OmitEmpty: true,
		}
	}

	fmt.Printf("%#v\n", keysInfo)
	statement, err := parseJenKeysAsStruct(name, keysInfo)
	if err != nil {
		return fmt.Errorf("Failed parsing dotted key: %s", err)
	}

	s.Add(statement)

	return nil
}
