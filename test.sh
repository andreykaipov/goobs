#!/bin/sh

cleanup() { docker stop obs-record obs-stream; }

setup() {
        echo "Setting up OBS instances for functional tests..."

        obs="$(docker container inspect -f '{{.State.Status}}' obs || true)"

        if [ "$obs" = running ]; then
                echo "Main OBS container is already running"
        else
                docker run --rm --detach --name obs -p 4444:1234 ghcr.io/andreykaipov/goobs
        fi

        # record and stream categories aren't totally idempotent so we need
        # a clean OBS docker instance each time

        echo "Spinning up OBS instances for 'record' and 'stream' tests"
        docker run --rm --detach --name obs-record -p 4445:1234 ghcr.io/andreykaipov/goobs
        docker run --rm --detach --name obs-stream -p 4446:1234 ghcr.io/andreykaipov/goobs

        covermode=count
        echo "mode: $covermode" >coverall.out
}

gotest() {
        category="$1"
        go test -v -run="^Test_$category$" -coverprofile=cover.out -coverpkg=./... -covermode=$covermode ./zz_generated.*test.go
        awk 'NR>1' cover.out >>coverall.out
}

main() {
        set -eu
        trap cleanup EXIT
        setup
        export OBS_PORT

        # note: `scenes` and `transitions` must be ran after `ui`
        #
        categories='
		config
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
                OBS_PORT=4444 gotest "$category"
        done

        OBS_PORT=4445 gotest record
        OBS_PORT=4446 gotest stream

        go tool cover -html coverall.out -o coverall.html
}

main "$@"
