package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"io/ioutil"
	"os/exec"
	"sort"
	"strings"

	. "github.com/dave/jennifer/jen"
	"golang.org/x/tools/go/packages"
)

var (
	goobs = "github.com/andreykaipov/goobs"
	root  = ""

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
		"outputs.",
		"scene_items.DuplicateSceneItem",
		"scene_items.GetGroupItemList",
		"scene_items.RemoveSceneItem",
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
	dirs, err := ioutil.ReadDir(root + "/api/requests")
	if err != nil {
		panic(err)
	}

	f := NewFile("goobs_test")
	f.HeaderComment("This file has been automatically generated. Don't edit it.")

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		category := dir.Name()
		path := fmt.Sprintf("%s/api/requests/%s", root, category)

		pkgs, err := packages.Load(&packages.Config{
			Mode: packages.NeedName | packages.NeedSyntax | packages.NeedImports,
			Dir:  path,
		})
		if err != nil {
			panic(err)
		}

		if len(pkgs) > 1 {
			panic(fmt.Errorf("really should only be one package in api/requests/%s ?", category))
		}

		pkg := pkgs[0]

		// Find structs
		structs := map[string]StructFieldMap{}

		for _, file := range pkg.Syntax {
			for objName, obj := range file.Scope.Objects {
				switch obj.Kind {
				case ast.Typ:
					switch typ := obj.Decl.(*ast.TypeSpec).Type.(type) {
					case *ast.StructType:
						structs[objName] = structFieldsToMap(typ, pkg)
					case *ast.Ident:
						continue
					default:
						continue
					}
				case ast.Con:
					continue
				default:
					continue
				}

			}
		}

		fmt.Println(category)

		test := generateRequestTest(pkg, category, structs)

		f.Add(test.Line())
	}

	if err := f.Save(fmt.Sprintf("%s/zz_generated._test.go", root)); err != nil {
		panic(err)
	}
}

func generateRequestTest(pkg *packages.Package, category string, structs map[string]StructFieldMap) *Statement {
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

	assert := "github.com/stretchr/testify/assert"

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
				s.List(Id("_"), Id("err")).Op("=").Id("client").Dot(pascal(category)).Dot(request).Call(Op("&").Qual(goobs+"/api/requests/"+category, request+"Params").Values(DictFunc(func(d Dict) {
					for field, fieldInfo := range structFields {
						fieldType := jenRenderUnsafe(fieldInfo.Type)

						d[Id(field)] = paramsMapper(request, field, fieldType)
					}
				})))
				s.Line()

				assertion := "NoError"
				for _, r := range requestsTestsAssertingErrors {
					if strings.HasPrefix(category+"."+request, r) {
						assertion = "Error"
						break
					}
				}
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
