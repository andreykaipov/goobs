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
	f.Render(buf)
	actual := buf.String()
	// fmt.Println(actual)

	fixture := strings.Join(strings.Split(t.Name(), "_")[2:], "_")
	expected, err := ioutil.ReadFile(fmt.Sprintf("testfixtures/dotparser/%s.go", fixture))
	if err != nil {
		panic(err)
	}

	assert.NoError(t, err)
	assert.Equal(t, string(expected), actual)
}

func Test_parseKeysAsJenStruct_basic_1(t *testing.T) {
	statement, err := parseJenKeysAsStruct("StartStreamingRequest", map[string]jen.Code{
		"stream":                   jen.Map(jen.String()).Interface(),
		"stream.type":              jen.Int(),
		"stream.metadata":          jen.Map(jen.String()).Interface(),
		"stream.settings":          jen.Map(jen.String()).Interface(),
		"stream.settings.server":   jen.String(),
		"stream.settings.key":      jen.String(),
		"stream.settings.use-auth": jen.Bool(),
		"stream.settings.username": jen.String(),
		"stream.settings.password": jen.String(),
	})

	assertJenStruct(t, statement, err)
}

func Test_parseKeysAsJenStruct_basic_2(t *testing.T) {
	statement, err := parseJenKeysAsStruct("IDK", map[string]jen.Code{
		"A.B": jen.String(),
		"C.D": jen.String(),
	})

	assertJenStruct(t, statement, err)
}

func Test_parseKeysAsJenStruct_basic_3(t *testing.T) {
	statement, err := parseJenKeysAsStruct("Ugh", map[string]jen.Code{
		"a":   jen.String(),
		"b":   jen.String(),
		"c.d": jen.Int(),
	}, map[string]string{
		"a":   "hello",
		"b":   "hi",
		"c.d": "bye",
	})

	assertJenStruct(t, statement, err)
}

func Test_parseKeysAsJenStruct_slices(t *testing.T) {
	statement, err := parseJenKeysAsStruct("GetSourcesListRequest", map[string]jen.Code{
		"Sources":          jen.Index().Map(jen.String()).Interface(),
		"Sources.*.Name":   jen.String(),
		"Sources.*.TypeId": jen.String(),
		"Sources.*.Type":   jen.String(),
	})

	assertJenStruct(t, statement, err)
}

func Test_parseKeysAsJenStruct_slices_nested(t *testing.T) {
	statement, err := parseJenKeysAsStruct("GetSourcesTypesListResponse", map[string]jen.Code{
		"Ids":                         jen.Index().Map(jen.String()).Interface(),
		"Ids.*.TypeId":                jen.String(),
		"Ids.*.DisplayName":           jen.String(),
		"Ids.*.Type":                  jen.String(),
		"Ids.*.DefaultSettings":       jen.Map(jen.String()).Interface(),
		"Ids.*.Caps":                  jen.Map(jen.String()).Interface(),
		"Ids.*.Caps.IsAsync":          jen.Bool(),
		"Ids.*.Caps.HasVideo":         jen.Bool(),
		"Ids.*.Caps.HasAudio":         jen.Bool(),
		"Ids.*.Caps.CanInteract":      jen.Bool(),
		"Ids.*.Caps.IsComposite":      jen.Bool(),
		"Ids.*.Caps.DoNotDuplicate":   jen.Bool(),
		"Ids.*.Caps.DoNotSelfMonitor": jen.Bool(),
		"Ids.*.Caps.Extra.*.Wow":      jen.Bool(),
		"Ids.*.Caps.Extra.*.Poopy":    jen.Bool(),
	})

	assertJenStruct(t, statement, err)
}

func Test_parseKeysAsJenStruct_slices_legacy(t *testing.T) {
	statement, err := parseJenKeysAsStruct("ReorderSceneItemsRequest_Legacy", map[string]jen.Code{
		"Items":        jen.Index().Map(jen.String()).Interface(),
		"Items[].Id":   jen.Int(),
		"Items[].Name": jen.String(),
	})

	assertJenStruct(t, statement, err)
}
