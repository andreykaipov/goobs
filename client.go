package goobs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/andreykaipov/goobs/api/events"
	"github.com/gorilla/websocket"
)

//go:generate internal/bin/funcopgen -type Client -prefix With -factory

type Client struct {
	Host string
	Port int

	IncomingEvents chan events.Event

	eventsConn *websocket.Conn // we use a different connection for listening to events
	conn       *websocket.Conn
	subclients
}

func (c *Client) Init() error {
	c.IncomingEvents = make(chan events.Event)

	conn, err := c.connect()
	if err != nil {
		return err
	}

	c.conn = conn
	setClients(c)
	return nil
}

func (c *Client) connect() (*websocket.Conn, error) {
	u := url.URL{
		Scheme: "ws",
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
	}

	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Client) Listen() {
	var err error

	c.eventsConn, err = c.connect()
	if err != nil {
		panic(err)
	}

	ch := make(chan json.RawMessage)
	go c.handleRawEvents(ch)
	c.handleUnknownEvents(ch)
}

func (c *Client) handleRawEvents(ch chan json.RawMessage) {
	for {
		raw := json.RawMessage{}
		if err := c.eventsConn.ReadJSON(&raw); err != nil {
			log.Fatal(err)
		}

		ch <- raw
	}
}

func (c *Client) handleUnknownEvents(ch chan json.RawMessage) {
	for raw := range ch {
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
