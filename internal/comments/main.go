package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

var (
	goobs = "github.com/andreykaipov/goobs"
	root  = ""

	version  = "4.9.1"
	comments = fmt.Sprintf("https://raw.githubusercontent.com/Palakis/obs-websocket/%s/docs/generated/comments.json", version)
)

func init() {
	var err error
	if root, err = run("git rev-parse --show-toplevel"); err != nil || root == "" {
		panic(err)
	}
}

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

	fmt.Println("Requests")
	generateRequests(data)

	fmt.Println("Events")
	generateEvents(data)

	fmt.Println("TypeDefs")
	generateTypeDefs(data)
}

func run(cmd string) (string, error) {
	output, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed running '%s': %s\n\n%s", cmd, err, output)
	}

	return strings.TrimSuffix(string(output), "\n"), nil
}
