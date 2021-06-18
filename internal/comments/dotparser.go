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
}

func parseJenKeysAsMap(lines map[string]keyInfo) (map[string]interface{}, error) {
	// e.g. keeps a reference from a.b to the corresponding struct
	mapReferences := map[string]map[string]interface{}{}

	for _, line := range sortedKeys(lines) {
		typ3 := lines[line].Type

		// prepend $. so the following loop always runs even for parts
		// with no dots, and replace [] as .* for legacy interpretations
		// of slices
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

	addTag := func(tag string) func(g *jen.Statement) {
		return func(s *jen.Statement) {
			s.Tag(map[string]string{"json": tag})
		}
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
		switch t := data.(type) {
		case map[string]interface{}:
			// if there's an *, redo the recursion with a slice
			for _, k := range sortedKeys(t) {
				v := t[k]
				if k == "*" {
					traverse(v, g, parent+" []")
					return
				}
			}

			// is a nested anon struct
			id := pascal(parent)
			idDeplural := id
			idType := id

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

				idType = "[]" + idDeplural
			}

			s := jen.Empty()
			traverseStruct(s, idDeplural, t)
			anonymousStructs = append(anonymousStructs, s)
			g.Id(id).Id(idType).Do(addTag(strings.TrimSuffix(parent, " []"))).Line()
			return
		case keyInfo:
			if t.Comment != "" {
				g.Comment(t.Comment)
			}
			g.Do(func(s *jen.Statement) {
				if t.Embedded {
					t.NoJSONTag = true
				} else {
					s.Id(pascal(parent))
				}
				s.Add(t.Type)
				if t.NoJSONTag {
					return
				}
				s.Do(addTag(parent))
			})
		default:
			panic("unhandled case idk")
		}

		g.Line()
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
