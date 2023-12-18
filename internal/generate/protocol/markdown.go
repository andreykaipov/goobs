package main

import (
	"fmt"
	"os"
	"strings"
)

func generateMarkdownTableMapping(protocol *Protocol, path string) {
	var sb strings.Builder

	sb.WriteString("# Request Mapping\n\n")
	sb.WriteString("Assuming we've defined a `goobs` client as follows:\n\n")
	sb.WriteString("```go\n")
	sb.WriteString(`client, err := goobs.New("localhost:4455", goobs.WithPassword("whatever"))`)
	sb.WriteString("\n```")
	sb.WriteString("\n\n")

	url := `https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#requests`
	sb.WriteString(fmt.Sprintf("The following tables show how to make the appropriate `goobs` request for any given any [obs-websocket request](%s):\n\n", url))

	generateMarkdownTableRequests(protocol, &sb)

	// the events are literally 1-1
	// kinda pointless to generate a table for them
	// generateMarkdownTableEvents(protocol, &sb)

	if err := os.WriteFile(path, []byte(sb.String()), 0644); err != nil {
		panic(err)
	}
}

func generateMarkdownTableRequests(protocol *Protocol, sb *strings.Builder) {
	mapping := map[string]map[string]string{}
	for _, request := range protocol.Requests {
		category := request.Category
		if _, ok := mapping[category]; !ok {
			mapping[category] = map[string]string{}
		}
		name := request.RequestType
		mapping[category][name] = fmt.Sprintf("client.%s.%s(...)", pascal(category), pascal(name))
	}

	for _, category := range sortedKeys(mapping) {
		obsToGoobs := mapping[category]

		sb.WriteString(fmt.Sprintf("## %s\n", category))
		sb.WriteString("| obs-websocket | goobs |\n")
		sb.WriteString("| --- | --- |\n")

		for _, obswebsocketRequest := range sortedKeys(obsToGoobs) {
			goobsRequest := obsToGoobs[obswebsocketRequest]
			sb.WriteString(fmt.Sprintf("| %s | `%s` |\n", obswebsocketRequest, goobsRequest))
		}
		sb.WriteString("\n")
	}
}

func generateMarkdownTableEvents(protocol *Protocol, sb *strings.Builder) {
	mapping := map[string]map[string]string{}
	for _, event := range protocol.Events {
		category := event.Category
		if _, ok := mapping[category]; !ok {
			mapping[category] = map[string]string{}
		}
		name := event.EventType
		mapping[category][name] = fmt.Sprintf("*events.%s", pascal(name))
	}

	sb.WriteString("## Events\n\n")
	for _, category := range sortedKeys(mapping) {
		obsToGoobs := mapping[category]

		sb.WriteString(fmt.Sprintf("### %s\n", category))
		sb.WriteString("| obs-websocket | goobs |\n")
		sb.WriteString("| --- | --- |\n")

		for _, obswebsocketEvent := range sortedKeys(obsToGoobs) {
			goobsEvent := obsToGoobs[obswebsocketEvent]
			sb.WriteString(fmt.Sprintf("| %s | `%s` |\n", obswebsocketEvent, goobsEvent))
		}
		sb.WriteString("\n")
	}
}
