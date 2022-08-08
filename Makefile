help:
	@echo "Please specify a task:"
	@awk -F: '/^[^\t#$$]+:[^=]+?$$/ {print "-",$$1}' Makefile

test: test.unit test.functional

test.unit:
	cd internal; go test -v -count 1 ./...
test.functional:
	./test.sh

generate:
	cd internal; go run ./generate/protocol/...
	cd internal; go run ./generate/tests/...
	$(MAKE) format

format:
	go install github.com/segmentio/golines@latest
	golines --shorten-comments --max-len=120 --reformat-tags --write-output *.go api

clean:
	find . -regextype awk -regex "./cover.+(out|html)" | xargs rm -f
	find . -regextype awk -regex "./.+_generated.[^.]+.go" | xargs rm -f
	find . -type d -empty -delete
	docker stop obs || true
