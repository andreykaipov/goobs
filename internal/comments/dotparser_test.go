package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/assert"
)

func assertJenStruct(t *testing.T, s *jen.Statement, err error) {
	buf := new(bytes.Buffer)
	f := jen.NewFile("testfixtures")
	f.Add(s)
	assert.NoError(t, f.Render(buf))
	actual := buf.String()
	// fmt.Println(actual)

	fixture := strings.Join(strings.Split(t.Name(), "_")[2:], "_")
	expected, err := ioutil.ReadFile(fmt.Sprintf("testfixtures/dotparser/%s/expected.go", fixture))
	if err != nil {
		panic(err)
	}

	assert.NoError(t, err)
	assert.Equal(t, string(expected), actual)
}

func Test_parseKeysAsJenStruct_basic_1(t *testing.T) {
	statement, err := parseJenKeysAsStruct("StartStreamingRequest", map[string]keyInfo{
		"stream":                   keyInfo{Type: jen.Map(jen.String()).Interface()},
		"stream.type":              keyInfo{Type: jen.Int()},
		"stream.metadata":          keyInfo{Type: jen.Map(jen.String()).Interface()},
		"stream.settings":          keyInfo{Type: jen.Map(jen.String()).Interface()},
		"stream.settings.server":   keyInfo{Type: jen.String()},
		"stream.settings.key":      keyInfo{Type: jen.String()},
		"stream.settings.use-auth": keyInfo{Type: jen.Bool()},
		"stream.settings.username": keyInfo{Type: jen.String()},
		"stream.settings.password": keyInfo{Type: jen.String()},
	})

	assertJenStruct(t, statement, err)
}

func Test_parseKeysAsJenStruct_basic_2(t *testing.T) {
	statement, err := parseJenKeysAsStruct("IDK", map[string]keyInfo{
		"A.B": keyInfo{Type: jen.String()},
		"C.D": keyInfo{Type: jen.String()},
	})

	assertJenStruct(t, statement, err)
}

func Test_parseKeysAsJenStruct_basic_3(t *testing.T) {
	statement, err := parseJenKeysAsStruct("Ugh", map[string]keyInfo{
		"a": keyInfo{
			Type:    jen.String(),
			Comment: "hello",
		},
		"b": keyInfo{
			Type:    jen.String(),
			Comment: "hi",
		},
		"c.d": keyInfo{
			Type:    jen.Int(),
			Comment: "bye",
		},
	})

	assertJenStruct(t, statement, err)
}

func Test_parseKeysAsJenStruct_slices(t *testing.T) {
	statement, err := parseJenKeysAsStruct("GetSourcesListRequest", map[string]keyInfo{
		"Sources":          keyInfo{Type: jen.Index().Map(jen.String()).Interface()},
		"Sources.*.Name":   keyInfo{Type: jen.String()},
		"Sources.*.TypeId": keyInfo{Type: jen.String()},
		"Sources.*.Type":   keyInfo{Type: jen.String()},
	})

	assertJenStruct(t, statement, err)
}

func Test_parseKeysAsJenStruct_slices_nested(t *testing.T) {
	statement, err := parseJenKeysAsStruct("GetSourcesTypesListResponse", map[string]keyInfo{
		"Ids":                         keyInfo{Type: jen.Index().Map(jen.String()).Interface()},
		"Ids.*.TypeId":                keyInfo{Type: jen.String()},
		"Ids.*.DisplayName":           keyInfo{Type: jen.String()},
		"Ids.*.Type":                  keyInfo{Type: jen.String()},
		"Ids.*.DefaultSettings":       keyInfo{Type: jen.Map(jen.String()).Interface()},
		"Ids.*.Caps":                  keyInfo{Type: jen.Map(jen.String()).Interface()},
		"Ids.*.Caps.IsAsync":          keyInfo{Type: jen.Bool()},
		"Ids.*.Caps.HasVideo":         keyInfo{Type: jen.Bool()},
		"Ids.*.Caps.HasAudio":         keyInfo{Type: jen.Bool()},
		"Ids.*.Caps.CanInteract":      keyInfo{Type: jen.Bool()},
		"Ids.*.Caps.IsComposite":      keyInfo{Type: jen.Bool()},
		"Ids.*.Caps.DoNotDuplicate":   keyInfo{Type: jen.Bool()},
		"Ids.*.Caps.DoNotSelfMonitor": keyInfo{Type: jen.Bool()},
		"Ids.*.Caps.Extra.*.Wow":      keyInfo{Type: jen.Bool()},
		"Ids.*.Caps.Extra.*.Poopy":    keyInfo{Type: jen.Bool()},
	})

	assertJenStruct(t, statement, err)
}

func Test_parseKeysAsJenStruct_slices_legacy(t *testing.T) {
	statement, err := parseJenKeysAsStruct("ReorderSceneItemsRequestLegacy", map[string]keyInfo{
		"Items":        keyInfo{Type: jen.Index().Map(jen.String()).Interface()},
		"Items[].Id":   keyInfo{Type: jen.Int()},
		"Items[].Name": keyInfo{Type: jen.String()},
	})

	assertJenStruct(t, statement, err)
}

func Test_parseKeysAsJenStruct_interfaces_1(t *testing.T) {
	statement, err := parseJenKeysAsStruct("Interfaces1", map[string]keyInfo{
		"a.my_interface": keyInfo{
			Type:      jen.Id("Interface"),
			NoJSONTag: true,
		},
	})

	assertJenStruct(t, statement, err)
}

func Test_parseKeysAsJenStruct_embedded_1(t *testing.T) {
	statement, err := parseJenKeysAsStruct("Embedded1", map[string]keyInfo{
		"a.b": keyInfo{
			Type:     jen.Id("EmbeddedDummy"),
			Embedded: true,
		},
	})

	assertJenStruct(t, statement, err)
}
