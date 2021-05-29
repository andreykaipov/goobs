package main

import (
	"fmt"
	"log"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/events"
	"github.com/andreykaipov/goobs/api/requests/sources"
)

var localhost = "172.24.80.1" // you'll likely have to change this :)

func main() {
	client, err := goobs.New(localhost+":4444", goobs.WithPassword("hello"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect()

	version, _ := client.General.GetVersion()
	fmt.Printf("Websocket server version: %s\n", version.ObsWebsocketVersion)
	fmt.Printf("OBS Studio version: %s\n", version.ObsStudioVersion)

	go func() {
		for event := range client.IncomingEvents {
			switch e := event.(type) {
			case *events.TransitionBegin:
				log.Printf("Transition the scene: %#v", e)
			case *events.Error:
				log.Printf("Got an error: %s", e.Err)
			default:
				log.Printf("Unhandled event: %#v", e.GetUpdateType())
			}
		}
	}()

	list, _ := client.Sources.GetSourcesList()

	fmt.Println("Available sources:")
	for _, v := range list.Sources {
		fmt.Printf("- %-25s of type %s\n", v.Name, v.TypeId)
	}

	update := func(volume float64) {
		da := "Desktop Audio"
		client.Sources.SetVolume(&sources.SetVolumeParams{
			Source: da,
			Volume: volume,
		})
		volumeResp, _ := client.Sources.GetVolume(&sources.GetVolumeParams{Source: da})
		fmt.Printf("Now %q is at %.3f\n", da, volumeResp.Volume)
	}

	update(0.3)
	update(0.5)
}
