package main

import (
	"fmt"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/sources"
)

var localhost = "172.18.144.1" // you'll likely have to change this :)

func main() {
	client := goobs.NewClient(goobs.WithHost(localhost), goobs.WithPort(4444))
	if _, err := client.Connect(); err != nil {
		panic(err)
	}
	defer client.Disconnect()

	version, _ := client.General.GetVersion()

	fmt.Printf("Websocket server version: %s\n", version.ObsWebsocketVersion)
	fmt.Printf("OBS Studio version: %s\n", version.ObsStudioVersion)

	sourcesResp, _ := client.Sources.GetSourcesList()

	fmt.Println("Available sources:")
	for _, v := range sourcesResp.Sources {
		resp, _ := client.Sources.GetVolume(&sources.GetVolumeParams{Source: v.Name})
		fmt.Printf("- %-25s %.3f\n", v.Name, resp.Volume)
	}

	da := "Desktop Audio"
	client.Sources.SetVolume(&sources.SetVolumeParams{
		Source: da,
		Volume: 0.25,
	})

	volumeResp, _ := client.Sources.GetVolume(&sources.GetVolumeParams{Source: da})
	fmt.Printf("Now %q is at %.3f\n", da, volumeResp.Volume)
}
