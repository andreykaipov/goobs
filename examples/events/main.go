package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andreykaipov/goobs"
)

type logger struct{}

func (l *logger) Printf(f string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, "\033[1;33m"+f+"\033[0m\n", v...)
}

func main() {
	client, err := goobs.New("localhost:4444", goobs.WithPassword("goodpassword"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect()

	version, _ := client.General.GetVersion()
	fmt.Printf("OBS Studio version: %s\n", version.ObsVersion)
	fmt.Printf("Websocket server version: %s\n", version.ObsWebSocketVersion)

	// This event loop is in the foreground now. If you mess around in OBS,
	// you'll see different events popping up!

	for event := range client.IncomingEvents {
		switch e := event.(type) {
		case error:
			log.Printf("Got an error: %s", e)
		default:
			log.Printf("Unhandled event: %[1]T: %#[1]v", e)
		}
	}
}
