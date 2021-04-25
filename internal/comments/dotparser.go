package main

import (
	"fmt"
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
type lineToType map[string]jen.Code

func f(lines lineToType) (map[string]interface{}, error) {
	// e.g. keeps a reference from a.b to the corresponding struct

	mapReferences := map[string]map[string]interface{}{}

	for line, typ3 := range lines {
		fmt.Println("doing", line)

		parts := strings.Split(line, ".")

		var m map[string]interface{}
		var ok bool

		for i, part := range parts[:len(parts)-1] {
			key := strings.Join(parts[0:i+1], ".")
			parentKey := strings.Join(parts[0:i], ".")

			if _, ok := mapReferences[parentKey][part].(*jen.Statement); ok {
				return nil, fmt.Errorf("1: tried to parse '%s' as '%T', but it already terminated at '%#v'", key, m, typ3)
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

	sorted := []string{}
	for k := range final {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	for _, k := range sorted {
		v := final[k]
		fmt.Printf("%s -> %#v\n", k, v)
	}

	return final, nil
}
