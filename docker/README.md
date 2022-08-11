## docker

This directory contains a Docker image with OBS and
[OBS-WebSocket](https://github.com/obsproject/obs-websocket/), allowing you to
run and control OBS headlessly.

While running OBS within a Docker container might not be particularly useful to
an end user since you'll have to jump through some hoops to get it to hook into
your USB devices and whatever, it's quite handy for testing obs-websocket client
libraries from the CI... and that's precisely how this library uses this image.
The obs-websocket server is bound to `localhost:1234` within the container.

One can even connect to it via VNC -- to observe exactly how your client is
interacting with it, or to manually setup some conditions for a particular test,
or for any other reason you'd like to log into it. The VNC server is bound to
`localhost:5900` within the container.

## usage

The most basic usage:

```console
❯ docker run --rm -it -p 4455:1234 ghcr.io/andreykaipov/goobs
```

Enable VNC:

```console
❯ docker run --rm -it -e vnc=1 -p 5900:5900 ghcr.io/andreykaipov/goobs
```

Optionally include a password:

```console
❯ docker run --rm -it -e vnc=1 -e vnc_password=delicious -p 5900:5900 ghcr.io/andreykaipov/goobs
```

Use some sort of VNC client (e.g. noVNC, VNC Viewer) to connect to it:

![Screenshot of a VNC client connected to the headless OBS Docker instance](vnc-example.png)

## configs

We can update the bundled scene and profile configs in this image by syncing the
configs to our local via `./sync.sh`, after modifying them from OBS itself while
we are VNC'd in.

A quick way to verify they were bundled correctly is by rerunning the rebuilt
image. Since copying the configs are last in our layers, this is rather quick:

```console
❯ docker run --rm -it --name obs -p 4455:1234 -e vnc=1 -p 5900:5900 "$(docker build -q docker)"
```
