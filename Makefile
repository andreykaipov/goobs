goobs_version                  := $(shell grep obs_version                    version.go | cut -d= -f2 | tr -dc '[0-9.]')
obs_websocket_protocol_version := $(shell grep obs_websocket_protocol_version version.go | cut -d= -f2 | tr -dc '[0-9.]')

export goobs_version
export obs_websocket_protocol_version

help:
	@echo "Please specify a task:"
	@awk -F: '/^[^.\t][^#$$]+:[^=]+?$$/ {print "-",$$1}' Makefile

.PHONY: test
test: test.unit test.functional
test.unit:
	cd internal; go test -v -count 1 ./...
test.functional:
	./test.sh

generate: generate.protocol generate.tests
	$(MAKE) format
generate.protocol:
	cd internal; go run ./generate/protocol/...
generate.tests:
	cd internal; go run ./generate/tests/...

format:
	go install github.com/segmentio/golines@latest
	golines --shorten-comments --max-len=120 --reformat-tags --write-output *.go api
	go install golang.org/x/tools/cmd/goimports@latest
	go list -f '{{.Dir}}' ./... | xargs goimports -w

clean:
	find . -regextype awk -regex "./cover.+(out|html)" | xargs rm -f
	find . -regextype awk -regex "./.+_generated\..+\.go" | xargs rm -f
	find . -type d -empty -delete
	docker stop obs || true

# 1. update version in version.go
# 2. commit all the changes
# 3. run this task
release:
	@if ! git diff --quiet; then \
		echo "Can't release... dirty worktree" ;\
		echo ;\
		git status ;\
		exit 1 ;\
	fi
	@gh release create v$(goobs_version) --generate-notes
