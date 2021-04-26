package goobs

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

//go:generate internal/bin/funcopgen -type Client -prefix With -factory

type Client struct {
	Host string
	Port int

	subclients
	conn *websocket.Conn
}

func (c *Client) Connect() (*Client, error) {
	u := url.URL{
		Scheme: "ws",
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
	}

	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	c.conn = conn
	setClients(c)

	return c, nil
}

func (c *Client) Disconnect() error {
	return c.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	)
}
