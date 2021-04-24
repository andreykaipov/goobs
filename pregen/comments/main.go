package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	for category, _ := range data.Requests {
		categorySnake := strings.ReplaceAll(category, " ", "_")
		categoryPascal := strings.ReplaceAll(strings.Title(category), " ", "")
		categoryClaustrophic := strings.ReplaceAll(category, " ", "")

		fmt.Printf("-- %s --\n", categoryPascal)

		f := NewFile(categoryClaustrophic)

		f.HeaderComment("This file has been automatically generated. Don't edit it.")

		f.HeaderComment("//go:generate go run github.com/andreykaipov/funcopgen -type Client -prefix With -factory -unexported") // lmao

		f.Comment(fmt.Sprintf("Client represents a client for '%s' requests", category))

		f.Add(
			Type().Id("Client").Struct(
				Id("conn").Op("*").Qual("github.com/gorilla/websocket", "Conn"),
			),
		)

		//fmt.Sprintf("%s/client.go", category)

		dir, err := run(fmt.Sprintf(`
			root="$(git rev-parse --show-toplevel)"
			dir="$root/pkg/%s"
			mkdir -p "$dir"
			printf "%%s" "$dir"
		`, categorySnake))
		if err != nil {
			panic(err)
		}

		if err := f.Save(fmt.Sprintf("%s/yy_generated.client.go", dir)); err != nil {
			panic(err)
		}
	}
}

func run(cmd string) (string, error) {
	output, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed running '%s': %s\n\n%s", cmd, err, output)
	}

	return string(output), nil
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
