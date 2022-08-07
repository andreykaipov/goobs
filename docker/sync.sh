#!/bin/sh

dcat() { docker exec -it obs cat "$@"; }

trap 'cd -' EXIT
cd config || exit

dcat /root/.config/obs-studio/global.ini | tr -d '\r' >./global.ini
dcat /root/.config/obs-studio/basic/profiles/docker/basic.ini | tr -d '\r' >./basic.ini
dcat /root/.config/obs-studio/basic/scenes/test.json | jq . >./scenes.json
