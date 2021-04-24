// This file has been automatically generated. Don't edit it.
//go:generate go run github.com/andreykaipov/funcopgen -type Client -prefix With -factory -unexported

package sources

import websocket "github.com/gorilla/websocket"

// Client represents a client for 'sources' requests
type Client struct {
	conn *websocket.Conn
}
