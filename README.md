# goobs

[![Go Reference](https://pkg.go.dev/badge/github.com/andreykaipov/goobs.svg)](https://pkg.go.dev/github.com/andreykaipov/goobs)

It's a Go client for
[Palakis/obs-websocket](https://github.com/Palakis/obs-websocket), allowing us
to interact with OBS Studio via Go.

## disclaimer

This project is still a work-in-progress.

It's currently using [v4.5.0 of the
protocol](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md).

## usage

Here's the contents of [examples/sources.go](examples/sources.go). For brevity,
we skip error checking:

```go
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
```

And the corresponding output:

```console
❯ go run examples/sources.go
2021/05/29 15:32:07 connecting to ws://172.24.80.1:4444
Websocket server version: 4.9.0
OBS Studio version: 26.1.1
Available sources:
- Video Capture Device      of type dshow_input
- Audio Output Capture      of type wasapi_output_capture
- Window Capture            of type window_capture
- Chat                      of type browser_source
- Mic/Aux                   of type wasapi_input_capture
- Desktop Audio             of type wasapi_output_capture
2021/05/29 15:32:07 Unhandled event: "SourceVolumeChanged"
Now "Desktop Audio" is at 0.300
2021/05/29 15:32:07 Unhandled event: "SourceVolumeChanged"
Now "Desktop Audio" is at 0.500
```
