
generate: clean
	go run ./pregen/comments/...
	go generate ./...

clean:
	find . \( -name 'yy_generated.*.go' -o -name 'zz_generated.*.go' \) -exec rm -f {} \;
	find . -type d -empty -delete
