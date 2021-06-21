# goobs

[![Go
Reference](https://pkg.go.dev/badge/github.com/andreykaipov/goobs.svg)](https://pkg.go.dev/github.com/andreykaipov/goobs)

It's a Go client for
[Palakis/obs-websocket](https://github.com/Palakis/obs-websocket), allowing us
to interact with OBS Studio via Go.

## installation

To add this client library to your module, simply `go get` it like any other Go
module after you've initialized your own:

```console
❯ go mod init blah
❯ go get github.com/andreykaipov/goobs
```

The latest release is v0.6.0 for [v4.9.1 of the obs-websocket
protocol](https://github.com/Palakis/obs-websocket/blob/4.9.1/docs/generated/protocol.md).

## usage

Usage is best demonstrated via a (hopefully self-explanatory) example. Here's
[examples/sources/main.go](examples/sources/main.go), showing off how we can
asynchronously listen for events, or send/receive synchronous
requests/responses, and how the two interact.

For brevity, error checks in a few places are omitted:

```go
package main

import ...

func main() {
	client, err := goobs.New(os.Getenv("WSL_HOST")+":4444", goobs.WithPassword("hello"))
	if err != nil {
		panic(err)
	}

	go func() {
		for event := range client.IncomingEvents {
			switch e := event.(type) {
			case *events.SourceVolumeChanged:
				fmt.Printf("Volume changed for %-25q: %f\n", e.SourceName, e.Volume)
			default:
				log.Printf("Unhandled event: %#v", e.GetUpdateType())
			}
		}
	}()

	fmt.Println("Setting random volumes for each source...")

	rand.Seed(time.Now().UnixNano())
	list, _ := client.Sources.GetSourcesList()

	for _, v := range list.Sources {
		if _, err := client.Sources.SetVolume(&sources.SetVolumeParams{
			Source: v.Name,
			Volume: rand.Float64(),
		}); err != nil {
			panic(err)
		}
	}

	fmt.Println("Test toggling the mute status of the first source...")

	name := list.Sources[0].Name
	resp, _ := client.Sources.GetVolume(&sources.GetVolumeParams{Source: name})
	fmt.Printf("%s is muted? %t\n", name, resp.Muted)

	_, _ = client.Sources.ToggleMute(&sources.ToggleMuteParams{Source: name})
	resp, _ = client.Sources.GetVolume(&sources.GetVolumeParams{Source: name})
	fmt.Printf("%s is muted? %t\n", name, resp.Muted)
}
```

And the corresponding output:

```console
❯ go run examples/sources/main.go
2021/06/14 02:22:18 connecting to ws://172.28.208.1:4444
Setting random volumes for each source...
Volume changed for "Chat"                   : 0.272767
Volume changed for "Window Capture"         : 0.791386
Volume changed for "Audio Output Capture"   : 0.777533
Volume changed for "Video Capture Device"   : 0.084827
Volume changed for "Mic/Aux"                : 0.104773
Volume changed for "Desktop Audio"          : 0.997565
Test toggling the mute status of the first source...
Chat is muted? false
2021/06/14 02:22:18 Unhandled event: "SourceMuteStateChanged"
Chat is muted? true
```

## development

This client is generated from the mess inside `./internal/comments`, reading the
[comments.json](https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/comments.json)
file found in [Palakis/obs-websocket](https://github.com/Palakis/obs-websocket).

Apart from `internal`, the only other hand-written code are the following files:

```console
❯ find *.go api/ -not -name '*_generated.*.go' -name '*.go'
client.go
api/requests/client.go
api/requests/requests.go
api/events/events.go
```

Iteration involves changing the generative code, running `make generate`, and
making sure the examples still work.
