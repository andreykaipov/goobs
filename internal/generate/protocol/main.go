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

	version  = "5.0.0"
	protocol = fmt.Sprintf("https://raw.githubusercontent.com/obsproject/obs-websocket/%s/docs/generated/protocol.json", version)
)

func init() {
	var err error
	if root, err = run("git rev-parse --show-toplevel"); err != nil || root == "" {
		panic(err)
	}
}

func main() {
	resp, err := http.Get(protocol)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	data := &Protocol{}
	if err := json.Unmarshal(body, data); err != nil {
		panic(err)
	}

	generateRequests(data.Requests)
	generateEvents(data.Events)
	generateEventSubscriptions(data.Enums, func(e *Enum) bool { return e.EnumType == "EventSubscription" })
	generateRequestStatuses(data.Enums, func(e *Enum) bool { return e.EnumType == "RequestStatus" })
}

func run(cmd string) (string, error) {
	output, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed running '%s': %s\n\n%s", cmd, err, output)
	}

	return strings.TrimSuffix(string(output), "\n"), nil
}
