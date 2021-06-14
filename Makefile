test:
	cd internal; go test -v -count 1 ./...

generate: clean
	cd internal; go run ./comments/
	cd internal; go build -o bin/ github.com/segmentio/golines
	./internal/bin/golines --shorten-comments --max-len=120 --reformat-tags --write-output api

clean:
	find . -name '*_generated.*.go' -exec rm -f {} \;
	find . -type d -empty -delete
	find . -type d -name bin -prune -exec rm -rf {} \;
