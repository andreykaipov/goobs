#!/bin/sh
# shellcheck disable=SC2086

set -eu

xvfb-run \
        --listen-tcp \
        --auth-file /var/run/Xauthority \
        --server-num 99 \
        --server-args '-screen 0 1200x800x24' \
        obs &

# indicates obs is up
until nc -vz localhost 1234; do
        echo "Waiting until obs-websocket has started successfully"
        sleep 1
done

if [ -z "${vnc-}" ]; then wait; fi

echo "Starting VNC server"

x11vncflags=
if [ -n "${vnc_password-}" ]; then
        x11vnc -storepasswd "$vnc_password" /tmp/vncpass
        x11vncflags='-rfbauth /tmp/vncpass'
fi

x11vnc \
        $x11vncflags \
        -rfbport 5900 \
        -display :99 \
        -forever \
        -auth /var/run/Xauthority
