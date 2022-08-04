#!/bin/sh

docker exec -it obs cat /root/.config/obs-studio/global.ini | tr -d '\r' >docker/global.ini
docker exec -it obs cat /root/.config/obs-studio/basic/profiles/docker/basic.ini | tr -d '\r' >docker/basic.ini
docker exec -it obs cat /root/.config/obs-studio/basic/scenes/docker.json | jq . >docker/scenes.json
