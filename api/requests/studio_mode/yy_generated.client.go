// This file has been automatically generated. Don't edit it.
//go:generate ../../../internal/bin/funcopgen -type Client -prefix With -factory -unexported

package studiomode

import websocket "github.com/gorilla/websocket"

// Client represents a client for 'studio mode' requests
type Client struct {
	conn *websocket.Conn
}
