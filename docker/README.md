## docker

This directory contains a Docker image with OBS and
[OBS-WebSocket](https://github.com/obsproject/obs-websocket/), allowing you to
run and control OBS headlessly.

While running OBS within a Docker container might not be particularly useful to
an end user since you'll have to jump through some hoops to get it to hook into
your USB devices and whatever, it's quite handy for testing OBS-Websocket client
libraries from the CI... and that's precisely how this library uses this image.

One can even connect to it via VNC, to observe exactly how your client is
interacting with it, or to manually setup some conditions for a paritcular test,
or for any other reason you'd like to log into it.

## usage

The most basic usage:

```console
❯ docker run --rm -it ghcr.io/andreykaipov/goobs
```

Enable VNC:

```console
❯ docker run --rm -it -e vnc=1 -p 5900:5900 ghcr.io/andreykaipov/goobs
```

Optionally include a password:

```console
❯ docker run --rm -it -e vnc=1 -e vnc_password=delicious -p 5900:5900 ghcr.io/andreykaipov/goobs
```

Use some sort of VNC client (e.g. noVNC, RealVNC) to connect to it:

![Screenshot of a VNC client connected to the headless OBS Docker instance](vnc-example2.png)
