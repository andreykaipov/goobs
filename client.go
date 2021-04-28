package goobs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/andreykaipov/goobs/api/events"
	eventsutil "github.com/andreykaipov/goobs/api/events/zz_util"
	"github.com/gorilla/websocket"
)

//go:generate internal/bin/funcopgen -type Client -prefix With -factory

type Client struct {
	Host string
	Port int

	IncomingEvents chan events.Event

	conn *websocket.Conn
	subclients
}

func (c *Client) connect() (*websocket.Conn, error) {
	u := url.URL{
		Scheme: "ws",
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
	}

	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	return conn, err
}

func (c *Client) Connect() (*Client, error) {
	conn, err := c.connect()
	if err != nil {
		return nil, err
	}

	c.conn = conn
	setClients(c)

	return c, nil
}

func (c *Client) Listen() {
	c.IncomingEvents = make(chan events.Event)

	// Use a separate connection for listening
	conn, err := c.connect()
	if err != nil {
		panic("Couldn't listen, bud")
	}

	for {
		raw := json.RawMessage{}
		if err := conn.ReadJSON(&raw); err != nil {
			log.Fatal(err)
		}

		unknownEvent := &events.EventCommon{}
		if err := json.Unmarshal(raw, unknownEvent); err != nil {
			log.Fatal(err)
		}

		eventType := unknownEvent.UpdateType

		knownEvent := eventsutil.GetEventForType(eventType)
		if knownEvent == nil {
			c.IncomingEvents <- unknownEvent
			continue
		}

		if err := json.Unmarshal(raw, knownEvent); err != nil {
			log.Fatal(err)
		}

		c.IncomingEvents <- knownEvent
	}
}

func (c *Client) Disconnect() error {
	return c.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	)
}
