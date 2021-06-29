default:
	@echo "Please specify a task:"
	@awk -F: '/^[^\t#.]+:[^=)]+?$$/ {print "-",$$1}' Makefile | tail -n+2 | sort

test: unit e2e

unit:
	cd internal; go test -v -count 1 ./...

.PHONY: e2e
e2e:
	@echo Setting up headless OBS for e2e...
	@docker stop goobs-e2e >/dev/null 2>&1 || true
	@docker run --rm --detach --name goobs-e2e -p 4545:4444 ghcr.io/andreykaipov/goobs/e2e >/dev/null
	@sleep 1 # lol
	go test -v -coverpkg ./... -coverprofile cover.out -v ./e2e/...
	go tool cover -html cover.out -o cover.html
	@docker stop goobs-e2e >/dev/null

generate: clean
	cd internal; go run ./comments/
	cd internal; go build -o bin/ github.com/segmentio/golines
	./internal/bin/golines --shorten-comments --max-len=120 --reformat-tags --write-output api

clean:
	find . -name '*_generated.*.go' -exec rm -f {} \;
	find . -type d -empty -delete
	find . -type d -name bin -prune -exec rm -rf {} \;
