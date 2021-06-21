package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dave/jennifer/jen"
)

type keyInfo struct {
	Type      jen.Code
	Comment   string
	NoJSONTag bool
	Embedded  bool
	OmitEmpty bool
}

func parseJenKeysAsMap(lines map[string]keyInfo) (map[string]interface{}, error) {
	// e.g. keeps a reference from a.b to the corresponding struct
	mapReferences := map[string]map[string]interface{}{}

	for _, line := range sortedKeys(lines) {
		typ3 := lines[line].Type

		// prepend $. so the following loop always runs even for parts
		// with no dots, and replace [] as .* to easily treat slices as
		// maps for now
		lineMod := "$." + strings.ReplaceAll(line, "[]", ".*")

		parts := strings.Split(lineMod, ".")

		var m map[string]interface{}
		var ok bool

		for i, part := range parts[:len(parts)-1] {
			key := strings.Join(parts[0:i+1], ".")
			parentKey := strings.Join(parts[0:i], ".")

			if val, ok := mapReferences[parentKey][part].(*jen.Statement); ok {
				// return nil, fmt.Errorf("1: tried to parse '%s' as '%T', but it already terminated at '%#v'", key, m, val)

				fmt.Fprintf(os.Stderr, "! Key %q already terminated at %T, but %q implies it's a map, so that's what we'll do\n", key, val, line)
			}

			if m, ok = mapReferences[key]; !ok {
				m = map[string]interface{}{}
				mapReferences[key] = m
			}

			if parentKey != "" {
				mapReferences[parentKey][part] = m
			}
		}

		lastPart := parts[len(parts)-1]

		if val, ok := m[lastPart]; ok {
			return nil, fmt.Errorf("2: wanted to terminate '%s' at '%#v', but it was already parsed as '%#T'", line, typ3, val)
		}

		m[lastPart] = lines[line]
	}

	final := map[string]interface{}{}
	for k, v := range mapReferences["$"] {
		if !strings.Contains(k, ".") {
			final[k] = v
		}
	}

	return final, nil
}

func parseJenKeysAsStruct(name string, lines map[string]keyInfo) (*jen.Statement, error) {
	m, err := parseJenKeysAsMap(lines)
	if err != nil {
		return nil, err
	}

	// mutually recrusive recursive with traverseStruct
	var traverse func(data interface{}, g *jen.Group, parent string)

	traverseStruct := func(s *jen.Statement, name string, t map[string]interface{}) {
		s.Id(name).StructFunc(func(subg *jen.Group) {
			for _, k := range sortedKeys(t) {
				v := t[k]
				traverse(v, subg, k)
			}
		}).Line()
	}

	// keep track of any anonymous structs as we'll want to make them
	// siblings with the original parent struct
	anonymousStructs := []*jen.Statement{}

	traverse = func(data interface{}, g *jen.Group, parent string) {
		var idType jen.Code
		id := pascal(parent)
		tag := strings.TrimSuffix(parent, "[]")

		switch t := data.(type) {
		case map[string]interface{}:
			// if there's an * in the keys, the parent key we're on
			// must have been a slice, so use "[]" as the marker,
			// and redo the recursion
			for _, k := range sortedKeys(t) {
				v := t[k]
				if k == "*" {
					traverse(v, g, parent+"[]")
					return
				}
			}

			// Since the type to our nested struct will always be
			// a pointer, it'll be omitted during encoding if that
			// value is ever nil, which seems always like the
			// preferred behavior. 🤷
			tag += ",omitempty"

			idType = jen.Op("*").Id(id) // is a nested anon struct, so use a pointer to it itself as the type
			idDeplural := id

			if strings.HasSuffix(id, "[]") {
				id = strings.TrimSuffix(id, "[]")
				idDeplural = id

				// try to depluralize lol
				// TODO add more exhaustive rules as necessary
				// because this isn't really robust
				if strings.HasSuffix(id, "s") {
					fmt.Printf("  ! %s is a slice and looks plural, we'll try to depluralize...\n", id)
					idDeplural = strings.TrimSuffix(id, "s")
				}

				idType = jen.Index().Op("*").Id(idDeplural)
			}

			s := jen.Empty()
			traverseStruct(s, idDeplural, t)
			anonymousStructs = append(anonymousStructs, s)
		case keyInfo:
			idType = t.Type

			if t.Comment != "" {
				g.Comment(t.Comment)
			}
			if t.Embedded {
				id = ""
				t.NoJSONTag = true
			}
			if t.OmitEmpty {
				tag += ",omitempty"
			}
			if t.NoJSONTag {
				tag = ""
			}
		default:
			panic("unhandled case idk")
		}

		g.Id(id).Add(idType).Do(func(s *jen.Statement) {
			if tag == "" {
				return
			}
			s.Tag(map[string]string{"json": tag})
		}).Line()
	}

	s := jen.Type()
	traverseStruct(s, name, m)
	for _, q := range anonymousStructs {
		s.Line().Type().Add(q).Line()
	}
	return s, nil
}

func pascal(text string) string {
	nodash := strings.ReplaceAll(text, "-", " ")
	noundies := strings.ReplaceAll(nodash, "_", " ")
	titled := strings.Title(noundies)
	return strings.ReplaceAll(titled, " ", "")
}
