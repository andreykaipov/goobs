#!/bin/sh
#
# This script bumps the obs_websocket_protocol_version in version.go to the next
# eligible version and generates the code for it.

set -eu

find_versions() {
        tags=$(
                gh api /repos/obsproject/obs-websocket/git/refs/tags -q .[].ref |
                        awk -F/ '/5[.][0-9]+[.][0-9]+/ {print $3}' |
                        grep -v - |
                        sort -Vr
        )

        current=$(grep ProtocolVersion version.go | tr -dc '0-9.')
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
}

compare_versions() {
        current_major=$(echo "$current" | cut -d. -f1)
        current_minor=$(echo "$current" | cut -d. -f2)
        current_patch=$(echo "$current" | cut -d. -f3)
        next_major=$(echo "$next" | cut -d. -f1)
        next_minor=$(echo "$next" | cut -d. -f2)
        next_patch=$(echo "$next" | cut -d. -f3)

        if [ "$next_major" -gt "$current_major" ]; then
                bump=major
        elif [ "$next_minor" -gt "$current_minor" ]; then
                bump=minor
        elif [ "$next_patch" -gt "$current_patch" ]; then
                bump='patch'
        else
                echo "$current is not greater than $next"
                echo "Nothing to update"
                exit 1
        fi
        echo "$bump" >/tmp/.goobs.protocol.bump
        echo "$next" >/tmp/.goobs.protocol.next
}

build_image() {
        if [ -z "$CI" ]; then return; fi
        gh workflow run image.yml
        sleep 3
        id=$(gh run list --workflow image.yml --event workflow_dispatch --json databaseId --jq 'first|.databaseId')
        gh run watch "$id" --interval 10
}

bump_versions() {
        sed -i "s/$current/$next/g" version.go README.md
        make generate
        ./script/update-readme-snippets.sh "$bump"
}

main() {
        find_versions
        compare_versions
        build_image
        bump_versions
}

main "$@"
