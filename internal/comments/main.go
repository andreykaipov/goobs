package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	. "github.com/dave/jennifer/jen"
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

	for category, _ := range data.Requests {
		fmt.Printf("-- %s --\n", category)

		categorySnake := strings.ReplaceAll(category, " ", "_")
		categoryPascal := strings.ReplaceAll(strings.Title(category), " ", "")
		categoryClaustrophic := strings.ReplaceAll(category, " ", "")

		// For the top-level client
		qualifier := "github.com/andreykaipov/goobs/api/" + categorySnake
		topClientFields = append(topClientFields, Id(categoryPascal).Op("*").Qual(qualifier, "Client"))
		topClientSetters = append(topClientSetters, Id("c").Dot(categoryPascal).Op("=").Qual(qualifier, "NewClient").Call(Qual(qualifier, "WithConn").Call(Id("c.conn"))))

		// Generate the category-level client
		f := NewFile(categoryClaustrophic)
		f.HeaderComment("This file has been automatically generated. Don't edit it.")
		f.HeaderComment("//go:generate go run github.com/andreykaipov/funcopgen -type Client -prefix With -factory -unexported") // lmao
		f.Comment(fmt.Sprintf("Client represents a client for '%s' requests", category))
		f.Add(Type().Id("Client").Struct(Id("conn").Op("*").Qual("github.com/gorilla/websocket", "Conn")))

		// Write the category-level client
		dir := fmt.Sprintf("%s/api/%s", root, categorySnake)
		if err := os.MkdirAll(dir, 0777); err != nil {
			panic(err)
		}

		if err := f.Save(fmt.Sprintf("%s/yy_generated.client.go", dir)); err != nil {
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

func run(cmd string) (string, error) {
	output, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed running '%s': %s\n\n%s", cmd, err, output)
	}

	return strings.TrimSuffix(string(output), "\n"), nil
}

// Comments the generated comments.json from Palakis
/*
type Comments struct {
	Events struct {
		General     []*Event `json:"general"`
		Media       []*Event `json:"media"`
		Other       []*Event `json:"other"`
		Profiles    []*Event `json:"profiles"`
		Recording   []*Event `json:"replay buffer"`
		SceneItems  []*Event `json:"scene items"`
		Scenes      []*Event `json:"scenes"`
		Sources     []*Event `json:"sources"`
		Streaming   []*Event `json:"streaming"`
		StudioMode  []*Event `json:"studio mode"`
		Transitions []*Event `json:"transitions"`
	} `json:"events"`
	Requests struct {
		General          []*Request `json:"general"`
		MediaControl     []*Request `json:"media control"`
		Outputs          []*Request `json:"outputs"`
		Profiles         []*Request `json:"profiles"`
		Recording        []*Request `json:"recording"`
		ReplayBuffer     []*Request `json:"replay buffer"`
		SceneCollections []*Request `json:"scene collections"`
		SceneItems       []*Request `json:"scene items"`
		Scenes           []*Request `json:"scenes"`
		Sources          []*Request `json:"sources"`
		Streaming        []*Request `json:"streaming"`
		StudioMode       []*Request `json:"studio mode"`
		Transitions      []*Request `json:"transitions"`
	} `json:"requests"`
	Typedefs []*TypeDef `json:"typedefs"`
}
*/
