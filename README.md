# goobs

[![Protocol Version][protocol-img]][protocol-url]
[![Documentation][doc-img]][doc-url]
[![Build Status][build-img]][build-url]
[![Go Report][goreport-img]][goreport-url]

[protocol-img]: https://img.shields.io/badge/obs--websocket-v4.9.1-blue?logo=obs-studio&style=flat-square
[protocol-url]: https://github.com/Palakis/obs-websocket/blob/4.9.1/docs/generated/protocol.md
[doc-img]: https://img.shields.io/badge/pkg.go.dev-reference-blue?logo=go&logoColor=white&style=flat-square
[doc-url]: https://pkg.go.dev/github.com/andreykaipov/goobs
[build-img]: https://img.shields.io/github/workflow/status/andreykaipov/goobs/test?logo=github&style=flat-square
[build-url]: https://github.com/andreykaipov/goobs/actions/workflows/ci.yml
[goreport-img]: https://goreportcard.com/badge/github.com/andreykaipov/goobs?logo=go&logoColor=white&style=flat-square
[goreport-url]: https://goreportcard.com/report/github.com/andreykaipov/goobs

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

## usage

Usage is best demonstrated through example! Here's
[examples/sources/main.go](examples/sources/main.go), showcasing both the
eventing and requests API. For brevity, error checks in a few places are
omitted:

```go
package main

import ...

func main() {
	client, err := goobs.New(
		os.Getenv("WSL_HOST")+":4444",
		goobs.WithPassword("hello"),                   // optional
		goobs.WithDebug(os.Getenv("OBS_DEBUG") != ""), // optional
	)
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

	if len(list.Sources) == 0 {
		fmt.Println("No sources!")
		os.Exit(0)
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

We can also run this example with `OBS_DEBUG` set to a non-empty string. If we
do, our client will log all of the raw sent requests and received responses.

For further examples, it might help browsing through
[`e2e/e2e_test.go`](./e2e/e2e_test.go), or through
[muesli/obs-cli](https://github.com/muesli/obs-cli) which consumes this library.

## development

The client library code is generated from the mess inside `./internal/comments`
by reading the generated [obs-websocket protocol
documentation](https://github.com/Palakis/obs-websocket/blob/4.9.1/docs/generated/comments.json).

Iteration typically involves changing the generative code, running `make
generate`, and a `make test`.
