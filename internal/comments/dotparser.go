package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
)

// a.b.c = int
// a.*.c = int
// a.b.*.c = int
//
// a = { b = { c = 1 } }
// a = [ { c = 1 }, { c = 2 } ]
// a = { b = [ { c = 1 }, { c = 2 } ] }
//
// a struct {
//   b struct {
//   	c int
//   }
// }
// a []struct {
//   c int
// }
// a struct {
//   b []struct {
//     c int
//   }
// }

// e.g. a.b.c = int

func parseJenKeysAsMap(lines map[string]jen.Code) (map[string]interface{}, error) {
	// e.g. keeps a reference from a.b to the corresponding struct

	mapReferences := map[string]map[string]interface{}{}

	for _, line := range sortedKeys(lines) {
		typ3 := lines[line]

		parts := strings.Split(line, ".")

		if len(parts) == 1 {
			fmt.Fprintf(os.Stderr, "! Key %q has no parts. Skipping...\n", line)
			continue
		}

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

		m[lastPart] = typ3
	}

	final := map[string]interface{}{}
	for k, v := range mapReferences {
		if !strings.Contains(k, ".") {
			final[k] = v
		}
	}

	return final, nil
}

func parseJenKeysAsStruct(name string, lines map[string]jen.Code) (*jen.Statement, error) {
	m, err := parseJenKeysAsMap(lines)
	if err != nil {
		return nil, err
	}

	var f func(data interface{}, g *jen.Group, parent string)

	f = func(data interface{}, g *jen.Group, parent string) {
		switch t := data.(type) {
		case map[string]interface{}:
			g.Id(parent).StructFunc(func(subg *jen.Group) {
				for _, k := range sortedKeys(t) {
					v := t[k]
					f(v, subg, k)
				}
			})
		case *jen.Statement:
			g.Add(jen.Id(parent).Add(t))
		default:
			panic("unhandled case idk")
		}
	}
	fmt.Printf("%#v\n", m)

	return jen.Type().CustomFunc(jen.Options{}, func(g *jen.Group) {
		f(m, g, name)
	}), nil
}

func sortedKeys(m interface{}) []string {
	var sorted []string
	i := 0

	switch t := m.(type) {
	case map[string]interface{}:
		sorted = make([]string, len(t))
		i := 0
		for k := range t {
			sorted[i] = k
			i++
		}
	case map[string]jen.Code:
		sorted = make([]string, len(t))
		for k := range t {
			sorted[i] = k
			i++
		}
	}

	sort.Strings(sorted)
	return sorted
}

/*
func printJSON(data interface{}, indent int) {
	switch t := data.(type) {
	case map[string]interface{}:
		fmt.Println("{")
		i := 0
		for k, v := range t {
			fmt.Printf("%*.s%q: ", indent+2, " ", k)
			printJSON(v, indent+2)
			if i < len(t)-2 {
				fmt.Printf(",\n")
			}
			i++
		}
		fmt.Printf("\n%*.s}", indent, " ")
	case *VariableArray:
		fmt.Println("[")
		for i, v := range t.Iterator() {
			fmt.Printf("%*.s", indent+2, " ")
			printJSON(v, indent+2)
			if i < t.Len()-1 {
				fmt.Printf(",\n")
			}
		}
		fmt.Printf("\n%*.s]", indent, " ")
	case nil:
		fmt.Print("null")
	case string:
		fmt.Printf("%q", data)
	default:
		fmt.Print(data)
	}
}
*/
