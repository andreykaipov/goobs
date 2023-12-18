## usage

todo

## advanced configuration

- `GOOBS_LOG` can be set to `trace`, `debug`, `info`, or `error` to better understand what our client is doing under the hood.

- `GOOBS_PROFILE` can be set to enable profiling.
  For example, the following will help us find unreleased memory:

  ```console
  ❯ GOOBS_PROFILE=memprofile=mem.out OBS_PORT=4455 go test -v -run=profile client_test.go
  ❯ go tool pprof -top -sample_index=inuse_space mem.out
  ```

  Set `GOOBS_PROFILE=help` to see all the other available options.
