#!/bin/sh

cleanup() {
        if [ -n "$CI" ]; then return; fi
        docker stop obs-record obs-stream
}

setup_docker() {
        echo "Setting up OBS instances for functional tests..."

        obs="$(docker container inspect -f '{{.State.Status}}' obs || true)"

        if [ "$obs" = running ]; then
                echo "Main OBS container is already running"
        else
                docker run --rm --detach --name obs -e vnc=1 -p 5900:5900 -p 4455:1234 ghcr.io/andreykaipov/goobs:latest
                sleep 3
        fi

        # record and stream categories aren't totally idempotent so we need
        # a clean OBS docker instance each time

        echo "Spinning up OBS instances for 'record' and 'stream' tests"
        docker run --rm --detach --name obs-record -p 4456:1234 ghcr.io/andreykaipov/goobs:latest
        docker run --rm --detach --name obs-stream -p 4457:1234 ghcr.io/andreykaipov/goobs:latest
}

setup() {
        if [ -z "$CI" ]; then setup_docker; fi
        covermode=count
        echo "mode: $covermode" >coverall.out
}

gotest() {
        export GOOBS_LOG=error
        category="$1"
        go test -v -run="^Test_$category$" -count 1 -coverprofile=cover.out -coverpkg=./... -covermode=$covermode ./...
        awk 'NR>1' cover.out >>coverall.out
}

main() {
        : "${CI=}"
        set -eu
        trap cleanup EXIT
        setup
        export OBS_PORT

        # note: `scenes` and `transitions` must be ran after `ui`
        #
        # (client|client_repeat|get_raw_response|multi_goroutine|profile)
        # config
        categories='
                filters
                general
                inputs
                mediainputs
                outputs
                rconfig
                sceneitems
                sources
                ui
                scenes
                transitions
        '

        for category in $categories; do
                echo "$category"
                OBS_PORT=4455 gotest "$category"
        done

        OBS_PORT=4456 gotest record
        OBS_PORT=4457 gotest stream

        go tool cover -html coverall.out -o coverall.html
}

main "$@"
