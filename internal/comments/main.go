package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strings"
	"unicode"

	. "github.com/dave/jennifer/jen"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	version  = "4.5.0"
	comments = fmt.Sprintf("https://raw.githubusercontent.com/Palakis/obs-websocket/%s/docs/generated/comments.json", version)
)

func main() {
	resp, err := http.Get(comments)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	data := &Comments{}
	if err := json.Unmarshal(body, data); err != nil {
		panic(err)
	}

	root, err := run("git rev-parse --show-toplevel")
	if err != nil {
		panic(err)
	}

	topClientFields := []Code{}
	topClientSetters := []Code{}

	// Sort the fields so we can traverse the map in a deterministic
	// order as we want the generated code to be the same between
	// subsequent runs.
	categories := make([]string, 0, len(data.Requests))
	for c := range data.Requests {
		categories = append(categories, c)
	}
	sort.Strings(categories)

	for _, category := range categories {
		fmt.Printf("-- %s --\n", category)

		categorySnake := strings.ReplaceAll(category, " ", "_")
		categoryPascal := strings.ReplaceAll(strings.Title(category), " ", "")
		categoryClaustrophic := strings.ReplaceAll(category, " ", "")

		// For the top-level client
		qualifier := "github.com/andreykaipov/goobs/api/" + categorySnake
		topClientFields = append(topClientFields, Id(categoryPascal).Op("*").Qual(qualifier, "Client"))
		topClientSetters = append(topClientSetters, Id("c").Dot(categoryPascal).Op("=").Qual(qualifier, "NewClient").Call(Qual(qualifier, "WithConn").Call(Id("c.conn"))))

		// Generate the category-level client
		client := NewFile(categoryClaustrophic)
		client.HeaderComment("This file has been automatically generated. Don't edit it.")
		client.HeaderComment("//go:generate ../../internal/bin/funcopgen -type Client -prefix With -factory -unexported") // lmao
		client.Comment(fmt.Sprintf("Client represents a client for '%s' requests", category))
		client.Add(Type().Id("Client").Struct(Id("conn").Op("*").Qual("github.com/gorilla/websocket", "Conn")))

		// Write the category-level client
		dir := fmt.Sprintf("%s/api/%s", root, categorySnake)
		if err := os.MkdirAll(dir, 0777); err != nil {
			panic(err)
		}
		if err := client.Save(fmt.Sprintf("%s/yy_generated.client.go", dir)); err != nil {
			panic(err)
		}

		// Generate the requests for the category
		requests := NewFile(categoryClaustrophic)
		requests.HeaderComment("This file has been automatically generated. Don't edit it.")
		for _, request := range data.Requests[category] {
			fmt.Printf("---- %s ----\n", request.Name)

			if request.Deprecated != "" {
				fmt.Fprintf(os.Stderr, "Request %q is deprecated\n", request.Name)
				continue
			}

			var structName string
			var structDef *Statement

			structName = request.Name + "Params"
			if structDef, err = parseParamsAsStruct(structName, request.Params); err != nil {
				panic(fmt.Errorf("Failed parsing 'Params' for request %q in category %q", request.Name, category))
			}
			requests.Add(structDef)

			structName = request.Name + "Response"
			if structDef, err = parseParamsAsStruct(structName, request.Returns); err != nil {
				panic(fmt.Errorf("Failed parsing 'Returns' for request %q in category %q", request.Name, category))
			}
			requests.Add(structDef)
		}
		if err := requests.Save(fmt.Sprintf("%s/yy_generated.requests.go", dir)); err != nil {
			panic(err)
		}
	}

	// Write utils for the top-level client
	f := NewFile("goobs")
	f.HeaderComment("This file has been automatically generated. Don't edit it.")
	f.Add(Type().Id("Subclients").Struct(topClientFields...))
	f.Add(Func().Id("setClients").Params(Id("c").Op("*").Id("Client")).Block(topClientSetters...))

	if err := f.Save(fmt.Sprintf("%s/yy_generated.client.go", root)); err != nil {
		panic(err)
	}
}

// parses "Params" or "Requests" fields from requests or event
func parseParamsAsStruct(name string, params []*Param) (*Statement, error) {
	keysToJenType := map[string]Code{}
	keysToComments := map[string]string{}

	for _, field := range params {
		fieldName, err := sanitizeText(field.Name)
		if err != nil {
			return nil, fmt.Errorf("Failed sanitizing field %#v", field)
		}

		var fieldType *Statement
		switch strings.Trim(strings.ReplaceAll(field.Type, "(optional)", ""), " ") {
		case "String":
			fieldType = String()
		case "int":
			fieldType = Int()
		case "Integer":
			fieldType = Int()
		case "double":
			fieldType = Float64()
		case "float":
			fieldType = Float64()
		case "bool":
			fieldType = Bool()
		case "boolean":
			fieldType = Bool()
		case "Boolean":
			fieldType = Bool()
		case "Object":
			fieldType = Map(String()).Interface()
		case "Array<String>":
			fieldType = Index().String()
		case "Array<Object>":
			fieldType = Index().Map(String()).Interface()
		case "Array<Source>":
			fieldType = Index().Map(String()).Interface()
		case "Array<Scene>":
			fieldType = Index().Map(String()).Interface()
		case "Scene|Array":
			fieldType = Index().Map(String()).Interface()
		default:
			panic(fmt.Errorf("%s is a weird type", field.Name))
		}

		keysToJenType[fieldName] = fieldType
		keysToComments[fieldName] = field.Description
	}

	statement, err := parseJenKeysAsStruct(name, keysToJenType, keysToComments)
	if err != nil {
		return nil, fmt.Errorf("Failed parsing dotted key: %s", err)
	}

	return statement, nil
}

func sanitizeText(text string) (string, error) {
	isMn := func(r rune) bool {
		// Mn: nonspacing marks
		return unicode.Is(unicode.Mn, r)
	}

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	clean, _, err := transform.String(t, text)
	return clean, err
}

func run(cmd string) (string, error) {
	output, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed running '%s': %s\n\n%s", cmd, err, output)
	}

	return strings.TrimSuffix(string(output), "\n"), nil
}
