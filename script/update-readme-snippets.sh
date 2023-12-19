#!/bin/sh
# shellcheck disable=SC2016
#
# This script updates the README.md file with the contents of the basic example.
# And it's outputs.

replace_markdown() {
        f="$1"
        markerprefix="$2"
        content=$(cat)
        awk -v content="$content" -v markerprefix="$markerprefix" '
                function replace() {
                        gsub(/\n"/, "\\n\"", content)
                        print "[//]: # (" markerprefix "-begin)"
                        print content
                        print "[//]: # (" markerprefix "-end)"
                }

                /('"$markerprefix"'-begin)/,/('"$markerprefix"'-end)/ {p=1; next}
                p { p=0; replace() }
                1
                END { if (p) replace() }
        ' "$f" >"$f.tmp"
        mv "$f.tmp" "$f"
}

# we need to run the basic example against the dummy OBS instance
ensure_obs() {
        echo "Setting up a dummy OBS instance..."

        obs="$(docker container inspect -f '{{.State.Status}}' obs || true)"

        if [ "$obs" = running ]; then
                echo "OBS container is already running"
        else
                docker run --rm --detach --name obs -p 4455:1234 ghcr.io/andreykaipov/goobs:latest
                sleep 3
        fi
}

main() {
        set -eu
        ensure_obs

        content=$(cat _examples/basic/main.go)
        printf '```go\n%s\n```' "$content" | replace_markdown README.md snippet-1

        draft_release=$(gh release list | awk -F'\t' '/Draft/ {print $3}' | tr -dc '0-9.' | head -n1)
        content=$(go run _examples/basic/main.go | sed "s/(devel)/$draft_release/")
        printf '```console\n%s\n```' "❯ go run _examples/basic/main.go\n$content" | replace_markdown README.md snippet-2
}

main "$@"
