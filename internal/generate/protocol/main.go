package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var (
	goobs = "github.com/andreykaipov/goobs"
	root  = ""

	version  = ""
	protocol = ""
)

func init() {
	log.SetFlags(0)

	var err error
	if root, err = run("git rev-parse --show-toplevel"); err != nil || root == "" {
		panic(err)
	}

	version = os.Getenv("obs_websocket_protocol_version")
	if version == "" {
		log.Fatal("Provide an OBS websocket version")
	}

	log.Printf("Generating with OBS websocket protocol version: %s", version)

	protocol = fmt.Sprintf("https://raw.githubusercontent.com/obsproject/obs-websocket/%s/docs/generated/protocol.json", version)
}

func main() {
	resp, err := http.Get(protocol)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	data := &Protocol{}
	if err := json.Unmarshal(body, data); err != nil {
		panic(err)
	}

	generateRequests(data.Requests)
	generateEvents(data.Events)
	generateEventSubscriptions(data.Enums, func(e *Enum) bool { return e.EnumType == "EventSubscription" })
	generateRequestStatuses(data.Enums, func(e *Enum) bool { return e.EnumType == "RequestStatus" })
	generateMarkdownTableMapping(data, fmt.Sprintf("%s/docs/request-mapping.md", root))
}

func run(cmd string) (string, error) {
	output, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed running '%s': %s\n\n%s", cmd, err, output)
	}

	return strings.TrimSuffix(string(output), "\n"), nil
}
