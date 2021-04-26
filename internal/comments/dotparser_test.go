package main

import (
	"bytes"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/assert"
)

func Test_parseKeysAsJenStruct_basic_1(t *testing.T) {
	statement, err := parseJenKeysAsStruct("StartStreamingRequest", map[string]jen.Code{
		"Stream":                   jen.Map(jen.String()).Interface(),
		"Stream.Type":              jen.Int(),
		"Stream.Metadata":          jen.Map(jen.String()).Interface(),
		"Stream.Settings":          jen.Map(jen.String()).Interface(),
		"Stream.Settings.Server":   jen.String(),
		"Stream.Settings.Key":      jen.String(),
		"Stream.Settings.UseAuth":  jen.Bool(),
		"Stream.Settings.Username": jen.String(),
		"Stream.Settings.Password": jen.String(),
	})

	buf := new(bytes.Buffer)
	statement.Render(buf)
	actual := buf.String()
	// fmt.Println(actual)

	expected := `type StartStreamingRequest struct {
	Stream struct {
		Metadata map[string]interface{}
		Settings struct {
			Key      string
			Password string
			Server   string
			UseAuth  bool
			Username string
		}
		Type int
	}
}`

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_parseKeysAsJenStruct_basic_2(t *testing.T) {
	statement, err := parseJenKeysAsStruct("IDK", map[string]jen.Code{
		"A.B": jen.String(),
		"C.D": jen.String(),
	})

	buf := new(bytes.Buffer)
	statement.Render(buf)
	actual := buf.String()
	// fmt.Println(actual)

	expected := `type IDK struct {
	A struct {
		B string
	}
	C struct {
		D string
	}
}`

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_parseKeysAsJenStruct_basic_3(t *testing.T) {
	statement, err := parseJenKeysAsStruct("Ugh", map[string]jen.Code{
		"a":   jen.String(),
		"b":   jen.String(),
		"c.d": jen.Int(),
	})

	buf := new(bytes.Buffer)
	statement.Render(buf)
	actual := buf.String()
	// fmt.Println(actual)

	expected := `type Ugh struct {
	a string
	b string
	c struct {
		d int
	}
}`

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_parseKeysAsJenStruct_slices(t *testing.T) {
	statement, err := parseJenKeysAsStruct("GetSourcesListRequest", map[string]jen.Code{
		"Sources":          jen.Index().Map(jen.String()).Interface(),
		"Sources.*.Name":   jen.String(),
		"Sources.*.TypeId": jen.String(),
		"Sources.*.Type":   jen.String(),
	})

	buf := new(bytes.Buffer)
	statement.Render(buf)
	actual := buf.String()
	// fmt.Println(actual)

	expected := `type GetSourcesListRequest struct {
	Sources []struct {
		Name   string
		Type   string
		TypeId string
	}
}`

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)

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

	buf := new(bytes.Buffer)
	statement.Render(buf)
	actual := buf.String()
	// fmt.Println(actual)

	expected := `type GetSourcesTypesListResponse struct {
	Ids []struct {
		Caps struct {
			CanInteract      bool
			DoNotDuplicate   bool
			DoNotSelfMonitor bool
			Extra            []struct {
				Poopy bool
				Wow   bool
			}
			HasAudio    bool
			HasVideo    bool
			IsAsync     bool
			IsComposite bool
		}
		DefaultSettings map[string]interface{}
		DisplayName     string
		Type            string
		TypeId          string
	}
}`

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_parseKeysAsJenStruct_slices_legacy(t *testing.T) {
	statement, err := parseJenKeysAsStruct("ReorderSceneItemsRequest_Legacy", map[string]jen.Code{
		"Items":        jen.Index().Map(jen.String()).Interface(),
		"Items[].Id":   jen.Int(),
		"Items[].Name": jen.String(),
	})

	buf := new(bytes.Buffer)
	statement.Render(buf)
	actual := buf.String()
	// fmt.Println(actual)

	expected := `type ReorderSceneItemsRequest_Legacy struct {
	Items []struct {
		Id   int
		Name string
	}
}`

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
