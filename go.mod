module github.com/andreykaipov/goobs

go 1.16

replace github.com/andreykaipov/funcopgen => ./internal/vendor/github.com/andreykaipov/funcopgen

require (
	github.com/andreykaipov/funcopgen v0.0.0 // indirect
	github.com/gorilla/websocket v1.4.2
)
