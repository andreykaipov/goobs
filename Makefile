
generate: clean
	cd internal; go run ./comments/...
	cd internal; go build -o bin/ github.com/andreykaipov/funcopgen
	go generate ./...

clean:
	find . \( -name 'yy_generated.*.go' -o -name 'zz_generated.*.go' \) -exec rm -f {} \;
	find . -type d -empty -delete
	find . -type d -name bin -prune -exec rm -rf {} \;
