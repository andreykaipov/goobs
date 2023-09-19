# goobs

[![Protocol Version][protocol-img]][protocol-url]
[![Documentation][doc-img]][doc-url]
[![Build Status][build-img]][build-url]
[![Go Report][goreport-img]][goreport-url]

[protocol-img]: https://img.shields.io/badge/obs--websocket-v5.1.0-blue?logo=obs-studio&style=flat-square
[protocol-url]: https://github.com/obsproject/obs-websocket/blob/5.1.0/docs/generated/protocol.md
[doc-img]: https://img.shields.io/badge/pkg.go.dev-reference-blue?logo=go&logoColor=white&style=flat-square
[doc-url]: https://pkg.go.dev/github.com/andreykaipov/goobs
[build-img]: https://img.shields.io/github/actions/workflow/status/andreykaipov/goobs/ci.yml?logo=github&style=flat-square&branch=master
[build-url]: https://github.com/andreykaipov/goobs/actions/workflows/ci.yml
[goreport-img]: https://goreportcard.com/badge/github.com/andreykaipov/goobs?logo=go&logoColor=white&style=flat-square
[goreport-url]: https://goreportcard.com/report/github.com/andreykaipov/goobs

It's a Go client for
[obsproject/obs-websocket](https://github.com/obsproject/obs-websocket),
allowing us to interact with OBS Studio from Go!

## installation

To add this library to your module, simply `go get` it like any other Go module
after you've initialized your own:

```console
❯ go mod init blah
❯ go get github.com/andreykaipov/goobs
```

## usage

Here's a basic example, where we grab the version and print out the scenes.
Check out the [examples](./examples) for more.

```go
package main

import (
	"fmt"
	"log"

	"github.com/andreykaipov/goobs"
)

func main() {
	client, err := goobs.New("localhost:4455", goobs.WithPassword("goodpassword"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect()

	version, err := client.General.GetVersion()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("OBS Studio version: %s\n", version.ObsVersion)
	fmt.Printf("Websocket server version: %s\n", version.ObsWebSocketVersion)

	resp, _ := client.Scenes.GetSceneList()
	for _, v := range resp.Scenes {
		fmt.Printf("%2d %s\n", v.SceneIndex, v.SceneName)
	}
}
```

This outputs the following:

```console
❯ go run examples/basic/main.go
OBS Studio version: 29.0.0
Websocket server version: 5.1.0
 1 Just Chatting
 2 Intermission
 3 Be Right Back 2
 4 Be Right Back
 5 Stream Starting Soon
 6 Background
 7 Camera Secondary
 8 Camera Primary
 9 Main 2
10 Main
```

## advanced configuration

- `GOOBS_LOG` can be set to `debug`, `info`, or `error` to better understand what our client is doing under the hood.

- `GOOBS_PROFILE` can be set to enable profiling.
  For example, the following will help us find unreleased memory:

  ```console
  ❯ GOOBS_PROFILE=memprofile=mem.out OBS_PORT=4455 go test -v -run=profile client_test.go
  ❯ go tool pprof -top -sample_index=inuse_space mem.out
  ```

  Set `GOOBS_PROFILE=help` to see all the other available options.
