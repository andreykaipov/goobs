# goobs

[![Protocol Version][protocol-img]][protocol-url]
[![Documentation][doc-img]][doc-url]
[![Build Status][build-img]][build-url]
[![Go Report][goreport-img]][goreport-url]

[protocol-img]: https://img.shields.io/badge/obs--websocket-v5.1.0-blue?logo=obs-studio&style=flat-square
[protocol-url]: https://github.com/obsproject/obs-websocket/blob/5.1.0/docs/generated/protocol.md
[doc-img]: https://img.shields.io/badge/pkg.go.dev-reference-blue?logo=go&logoColor=white&style=flat-square
[doc-url]: https://pkg.go.dev/github.com/andreykaipov/goobs
[build-img]: https://img.shields.io/github/actions/workflow/status/andreykaipov/goobs/ci.yml?logo=github&style=flat-square&branch=main
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

Here's the [basic example](./_examples/basic), where we connect to the server and print out some versions.
Check out the [other examples](./_examples) for more.

[//]: # (snippet-1-begin)
blah
[//]: # (snippet-1-end)

This outputs the following:

[//]: # (snippet-2-begin)
blah
[//]: # (snippet-2-end)

## advanced configuration

- `GOOBS_LOG` can be set to `trace`, `debug`, `info`, or `error` to better understand what our client is doing under the hood.

- `GOOBS_PROFILE` can be set to enable profiling.
  For example, the following will help us find unreleased memory:

  ```console
  ❯ GOOBS_PROFILE=memprofile=mem.out OBS_PORT=4455 go test -v -run=profile client_test.go
  ❯ go tool pprof -top -sample_index=inuse_space mem.out
  ```

  Set `GOOBS_PROFILE=help` to see all the other available options.
