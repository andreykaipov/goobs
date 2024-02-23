# goobs

[![Protocol Version][protocol-img]][protocol-url]
[![Documentation][doc-img]][doc-url]
[![Build Status][build-img]][build-url]
[![Go Report][goreport-img]][goreport-url]

[protocol-img]: https://img.shields.io/badge/obs--websocket-v5.4.2-blue?logo=obs-studio&style=flat-square
[protocol-url]: https://github.com/obsproject/obs-websocket/blob/5.4.2/docs/generated/protocol.md
[doc-img]: https://img.shields.io/badge/pkg.go.dev-reference-blue?logo=go&logoColor=white&style=flat-square
[doc-url]: https://pkg.go.dev/github.com/andreykaipov/goobs
[build-img]: https://img.shields.io/github/actions/workflow/status/andreykaipov/goobs/ci.yml?logo=github&style=flat-square&branch=main
[build-url]: https://github.com/andreykaipov/goobs/actions/workflows/ci.yml
[goreport-img]: https://goreportcard.com/badge/github.com/andreykaipov/goobs?logo=go&logoColor=white&style=flat-square
[goreport-url]: https://goreportcard.com/report/github.com/andreykaipov/goobs

Interact with OBS Studio from Go!

## installation

To use this library in your project, add it as a module after you've initialized your own:

```console
❯ go mod init github.com/beautifulperson/my-cool-obs-thing
❯ go get github.com/andreykaipov/goobs
```

## usage

The following example connects to the server and prints out some versions.

Check out the [docs](./docs/README.md) for more info, or just jump right into the [other examples](./_examples)!

[//]: # (snippet-1-begin)
```go
package main

import (
	"fmt"

	"github.com/andreykaipov/goobs"
)

func main() {
	client, err := goobs.New("localhost:4455", goobs.WithPassword("goodpassword"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect()

	version, err := client.General.GetVersion()
	if err != nil {
		panic(err)
	}

	fmt.Printf("OBS Studio version: %s\n", version.ObsVersion)
	fmt.Printf("Server protocol version: %s\n", version.ObsWebSocketVersion)
	fmt.Printf("Client protocol version: %s\n", goobs.ProtocolVersion)
	fmt.Printf("Client library version: %s\n", goobs.LibraryVersion)
}
```
[//]: # (snippet-1-end)

The corresponding output:

[//]: # (snippet-2-begin)
```console
❯ go run _examples/basic/main.go
OBS Studio version: 30.1.0
Server protocol version: 5.4.2
Client protocol version: 5.4.2
Client library version: 1.2.2
```
[//]: # (snippet-2-end)
