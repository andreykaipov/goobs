// This file has been automatically generated. Don't edit it.

package transitions

import websocket "github.com/gorilla/websocket"

type Option func(*Client)

func NewClient(opts ...Option) *Client {
	o := &Client{}

	for _, opt := range opts {
		opt(o)
	}

	return o
}

func WithConn(x *websocket.Conn) Option {
	return func(o *Client) {
		o.conn = x
	}
}
