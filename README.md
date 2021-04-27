# goobs

[![Go Reference](https://pkg.go.dev/badge/github.com/andreykaipov/goobs.svg)](https://pkg.go.dev/github.com/andreykaipov/goobs)

It's a Go client for
[Palakis/obs-websocket](https://github.com/Palakis/obs-websocket), allowing us
to interact with OBS Studio via Go.

## usage

Here's the contents of [examples/sources.go](examples/sources.go). For brevity,
we skip error checking:

```go
package main

import (
	"fmt"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/sources"
)

func main() {
	client := goobs.NewClient(goobs.WithHost("127.0.0.1"), goobs.WithPort(4444))
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
```

And the corresponding output:

```console
❯ go run examples/sources.go
2021/04/27 00:17:33 connecting to ws://127.0.0.1:4444
Websocket server version: 4.9.0
OBS Studio version: 26.1.1
Available sources:
- Chat                      1.000
- Window Capture            1.000
- Audio Output Capture      1.000
- Video Capture Device      1.000
- Mic/Aux                   1.000
- Desktop Audio             1.000
Now "Desktop Audio" is at 0.250
```
