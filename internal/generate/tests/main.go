package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"reflect"
	"sort"
	"strings"

	. "github.com/dave/jennifer/jen"
)

var (
	goobs     = "github.com/andreykaipov/goobs"
	root      = ""
	assert    = "github.com/stretchr/testify/assert"
	websocket = "github.com/gorilla/websocket"

	// These are requests that should be tested with assert.Error instead of
	// assert.NoError. They are either very hard to generate idempotently or
	// I am just too lazy to get them to not error.
	//
	requestsTestsAssertingErrors = []string{
		"config.CreateSceneCollection",               // we start with a `SceneCollectionName` collection already
		"filters.SetSourceFilterName",                // not idempotent
		"general.CallVendorRequest",                  // no other third party plugins in my obs image
		"general.Sleep",                              // only available in request batches and i don't wanna do that
		"inputs.GetInputPropertiesListPropertyItems", // idk what properties are
		"inputs.PressInputPropertiesButton",          // idk what properties are
		"inputs.SetInputAudioMonitorType",            // audio monitoring not available on this platform
		"inputs.SetInputName",                        // not idempotent
		"inputs.SetInputVolume",                      // too tricky to remove one of the two volume fields
		"!outputs.GetOutputList",
		"outputs.",
		"sceneitems.DuplicateSceneItem",
		"sceneitems.GetGroupSceneItemList",
		"sceneitems.RemoveSceneItem",
		"scenes.CreateScene",
		"scenes.RemoveScene",
		"scenes.SetSceneName",
		"stream.SendStreamCaption",
		"stream.StopStream",
		"transitions.SetCurrentSceneTransitionSettings", // seems like no default transition has custom settings
		"ui.OpenInputInteractDialog",                    // I think only the browser source supported the interact dialog

		// not idempotent between runs at all so these also need to be
		// run at the end of all others
		"record.ResumeRecord",
		"record.StopRecord",
	}
)

func init() {
	var err error
	if root, err = run("git rev-parse --show-toplevel"); err != nil || root == "" {
		panic(err)
	}
}

func run(cmd string) (string, error) {
	output, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed running '%s': %s\n\n%s", cmd, err, output)
	}

	return strings.TrimSuffix(string(output), "\n"), nil
}

// renders our statement. if it errors because of formatting issues, parse the
// invalid Go code from the error. don't want to maintain a fork for the proper
// solution
//
// see https://github.com/dave/jennifer/pull/94
func jenRenderUnsafe(s *Statement) string {
	buf := &bytes.Buffer{}
	if err := s.Render(buf); err != nil {
		return strings.Join(strings.Split(err.Error(), "\n")[1:], "\n")
	}
	return buf.String()
}

func main() {
	f := NewFile("goobs_test")
	f.HeaderComment("This file has been automatically generated. Don't edit it.")

	f.Add(generateClientTest().Line())

	// find subclients at the root of goobs, using that to then find the
	// structs within each of the subclient categories

	structs, err := findStructs(root)
	if err != nil {
		panic(err)
	}

	subclients := sortedKeys(structs["subclients"])

	for _, subclient := range subclients {
		category := strings.ToLower(subclient)

		structs, err := findStructs(root + "/api/requests/" + category)
		if err != nil {
			panic(err)
		}

		test := generateRequestTest(subclient, category, structs)

		f.Add(test.Line())
	}

	if err := f.Save(fmt.Sprintf("%s/zz_generated._test.go", root)); err != nil {
		panic(err)
	}
}

func generateClientTest() *Statement {
	s := Line()

	s.Func().Id("Test_client").Params(Id("t").Op("*").Qual("testing", "T")).Block(
		Var().Id("err").Error(),
		// tries to connect incorrectly
		List(Id("_"), Id("err")).Op("=").Qual(goobs, "New").Call(
			Lit("localhost:").Op("+").Qual("os", "Getenv").Call(Lit("OBS_PORT")),
			Qual(goobs, "WithPassword").Call(Lit("wrongpassword")),
			Qual(goobs, "WithRequestHeader").Call(Qual("net/http", "Header").Values(Dict{
				Lit("User-Agent"): Index().String().Values(Lit("goobs-e2e/0.0.0")),
			})),
		),
		Qual(assert, "Error").Call(Id("t"), Id("err")),
		Qual(assert, "IsType").Call(Id("t"), Op("&").Qual(websocket, "CloseError").Block(), Id("err")),
		Qual(assert, "Equal").Call(Id("t"), Id("err").Assert(Op("*").Qual(websocket, "CloseError")).Dot("Code"), Lit(4009)),

		// tries to connect to a nonrunning server
		List(Id("_"), Id("err")).Op("=").Qual(goobs, "New").Call(
			Lit("localhost:42069"),
			Qual(goobs, "WithPassword").Call(Lit("wrongpassword")),
			Qual(goobs, "WithRequestHeader").Call(Qual("net/http", "Header").Values(Dict{
				Lit("User-Agent"): Index().String().Values(Lit("goobs-e2e/0.0.0")),
			})),
		),
		Qual(assert, "Error").Call(Id("t"), Id("err")),
		Qual(assert, "IsType").Call(Id("t"), Op("&").Qual("net", "OpError").Block(), Id("err")),
	)

	return s
}

func generateRequestTest(subclient, category string, structs map[string]StructFieldMap) *Statement {
	paramsMapper := func(request, field, fieldType string) *Statement {
		var val *Statement
		switch fieldType {
		case "string":
			lit := ""
			switch field {
			case "Realm":
				lit = "OBS_WEBSOCKET_DATA_REALM_GLOBAL"
			case "MediaAction":
				lit = "OBS_WEBSOCKET_MEDIA_INPUT_ACTION_PAUSE"
			case "SceneItemBlendMode":
				lit = "OBS_BLEND_NORMAL"
			case "KeyId":
				lit = "OBS_KEY_SHIFT"
			case "HotkeyName":
				lit = "OBSBasic.ShowContextBar"
			case "StreamServiceType":
				lit = "rtmp_custom"
			case "FilterKind":
				lit = "scroll_filter"
			case "InputKind":
				lit = "ffmpeg_source"
			case "SceneName":
				switch request {
				case "RemoveScene":
					lit = "nonexistent"
				case "RemoveSceneItem":
					lit = "nonexistent"
				default:
					lit = "Scene"
				}
			case "NewSceneName":
				lit = "Scene"
			case "TransitionName":
				lit = "Cut"
			case "ImageFormat":
				lit = "png"
			case "VideoMixType":
				lit = "OBS_WEBSOCKET_VIDEO_MIX_TYPE_PREVIEW"
			case "ProjectorGeometry":
				lit = ""
			default:
				switch category {
				case "inputs":
					lit = "test2"
				default:
					lit = "test"
				}
			}
			val = Lit(lit)
		case "float64":
			var n float64
			switch field {
			case "TransitionDuration":
				n = 50.0
			default:
				switch category {
				case "config":
					n = 10.0
				case "sources":
					n = 8.0
				default:
					n = 1.0
				}
			}
			val = Lit(n)
		case "*bool":
			val = Op("&").Index().Bool().Values(True()).Index(Lit(0))
		case "interface{}":
			switch field {
			case "EventData":
				val = Map(String()).Bool().Values(Dict{Lit("test"): True()})
			default:
				val = Lit("")
			}
		case "map[string] interface{}":
			val = Map(String()).Interface().Values(Dict{
				Lit("test"): Lit("test"),
			})
		case "*typedefs.StreamServiceSettings":
			val = Op("&").Qual(goobs+"/api/typedefs", "StreamServiceSettings").Values(Dict{})
		case "*typedefs.KeyModifiers":
			val = Op("&").Qual(goobs+"/api/typedefs", "KeyModifiers").Values()
		case "*typedefs.InputAudioTracks":
			val = Op("&").Qual(goobs+"/api/typedefs", "InputAudioTracks").Values(Dict{
				Lit("test"): True(),
			})
		case "*typedefs.SceneItemTransform":
			val = Op("&").Qual(goobs+"/api/typedefs", "SceneItemTransform").Values(Dict{
				Id("BoundsType"):   Lit("OBS_BOUNDS_NONE"),
				Id("BoundsWidth"):  Lit(1.0),
				Id("BoundsHeight"): Lit(1.0),
			})
		default:
			panic(fmt.Errorf("%s.%s: param field %s: unhandled type: %q", category, request, field, fieldType))
		}

		return val
	}

	return Func().Id("Test_"+category).Params(Id("t").Op("*").Qual("testing", "T")).Block(
		List(Id("client"), Id("err")).Op(":=").Qual(goobs, "New").Call(
			Lit("localhost:").Op("+").Qual("os", "Getenv").Call(Lit("OBS_PORT")),
			Qual(goobs, "WithPassword").Call(Lit("goodpassword")),
			Qual(goobs, "WithRequestHeader").Call(Qual("net/http", "Header").Values(Dict{
				Lit("User-Agent"): Index().String().Values(Lit("goobs-e2e/0.0.0")),
			})),
		),
		Qual(assert, "NoError").Call(Id("t"), Id("err")),
		//Defer().Id("client").Dot("Disconnect").Call(),

		Id("t").Dot("Cleanup").Call(Func().Call().Block(
			Id("client").Dot("Disconnect").Call(),
		)),

		Do(func(s *Statement) {
			structNames := customSortedKeys(structs)
			for _, name := range structNames {
				structFields := structs[name]

				if !strings.HasSuffix(name, "Params") {
					continue
				}

				request := strings.TrimSuffix(name, "Params")
				// fmt.Printf("%s %s\n", category, request)

				s.Line()
				s.List(Id("_"), Id("err")).Op("=").Id("client").Dot(subclient).Dot(request).Call(Op("&").Qual(goobs+"/api/requests/"+category, request+"Params").Values(DictFunc(func(d Dict) {
					for field, fieldInfo := range structFields {
						fieldType := jenRenderUnsafe(fieldInfo.Type)

						d[Id(field)] = paramsMapper(request, field, fieldType)
					}
				})))
				s.Line()

				assertion := "NoError"
				for _, r := range requestsTestsAssertingErrors {
					prefix := category + "." + request
					// make sure to put negations before
					// a wider prefix pattern
					if strings.HasPrefix("!"+prefix, r) {
						break
					}
					if strings.HasPrefix(prefix, r) {
						assertion = "Error"
						break
					}
				}

				// stop record is right after start record, and
				// this causes flaky tests in GH actions
				// (sometimes the output starts and sometimes it
				// doesn't in time)
				// just don't assert for anything
				if request == "StopRecord" {
					s.Id("t").Dot("Logf").Call(Lit("skipped: %s"), Id("err"))
					continue
				}

				s.If(Id("err").Op("!=").Nil()).Block(
					Id("t").Dot("Logf").Call(Lit("%s"), Id("err")),
				)
				s.Line()
				s.Qual(assert, assertion).Call(Id("t"), Id("err"))
			}
		}),
	)
}

func pascal(text string) string {
	nodash := strings.ReplaceAll(text, "-", " ")
	noundies := strings.ReplaceAll(nodash, "_", " ")
	titled := strings.Title(noundies)
	return strings.ReplaceAll(titled, " ", "")
}

func customSortedKeys(m map[string]StructFieldMap) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}

	isRemove := func(s string) bool { return strings.HasPrefix(s, "Remove") }

	// prefer RemoveX requests at the end to make idempotency easier
	//
	sort.Slice(keys, func(i, j int) bool {
		a := keys[i]
		b := keys[j]

		if isRemove(a) && !isRemove(b) {
			return false
		}
		if !isRemove(a) && isRemove(b) {
			return true
		}

		return a < b
	})

	return keys
}

func sortedKeys(m interface{}) []string {
	keys := reflect.ValueOf(m).MapKeys()

	sorted := make([]string, len(keys))
	i := 0
	for _, key := range keys {
		sorted[i] = key.Interface().(string)
		i++
	}
	sort.Strings(sorted)

	return sorted
}
