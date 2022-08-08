package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/events"
)

func main() {
	client, _ := goobs.New("localhost:4444", goobs.WithPassword("goodpassword"))
	defer client.Disconnect()

	version, _ := client.General.GetVersion()
	fmt.Printf("OBS Studio version: %s\n", version.ObsVersion)
	fmt.Printf("Websocket server version: %s\n", version.ObsWebSocketVersion)

	for event := range client.IncomingEvents {
		switch e := event.(type) {
		case *events.ExitStarted:
			log.Printf("Server closing; gracefully shutting down")
			log.Printf("Bye")
			os.Exit(0)
		case error:
			log.Printf("Error: %s", e)
		default:
			log.Printf("Unhandled event: %#v", e)
		}
	}
}
