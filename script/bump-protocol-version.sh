#!/bin/sh
#
# This script bumps the obs_websocket_protocol_version in version.go to the next
# eligible version and generates the code for it.

set -eu

tags=$(
        gh api /repos/obsproject/obs-websocket/git/refs/tags -q .[].ref |
                awk -F/ '/5[.][0-9]+[.][0-9]+/ {print $3}' |
                grep -v - |
                sort -Vr
)

current=$(grep obs_websocket_protocol_version version.go | tr -dc '0-9.')
next=$(echo "$tags" | sort -V | awk -v current="$current" '$0==current{getline;print;exit}')
latest=$(echo "$tags" | sort -V | tail -n1)

cat >&2 <<EOF
Current: $current
Next:    $next
Latest:  $latest
EOF

version_regex='^[0-9]+\.[0-9]+\.[0-9]+$'

for v in "$current" "$next" "$latest"; do
        if ! echo "$v" | grep -qE "$version_regex"; then
                echo "Doesn't look like a version: $v"
                exit 1
        fi
done

sed -i "s/$current/$next/g" version.go
make generate
