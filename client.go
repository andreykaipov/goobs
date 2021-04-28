package goobs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/andreykaipov/goobs/api/events"
	"github.com/gorilla/websocket"
)

// Client represents a client to an OBS websockets server.
type Client struct {
	IncomingEvents chan events.Event

	host string
	subclients

	// we use two different connections for events and requests to avoid
	// race conditions when reading events and request responses
	eventsConn   *websocket.Conn
	requestsConn *websocket.Conn
}

// Option represents a functional option of a Client.
type Option func(*Client)

/*
New creates and configures a client to interact with the OBS websockets server.
It also opens up a connection, so be sure to check the error.
*/
func New(host string, opts ...Option) (c *Client, err error) {
	c = &Client{
		host:           host,
		IncomingEvents: make(chan events.Event),
	}

	for _, opt := range opts {
		opt(c)
	}

	c.requestsConn, err = c.connect()
	if err != nil {
		return nil, err
	}

	setClients(c)

	return c, nil
}

func (c *Client) connect() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: c.host}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

/*
Listen starts listening for events from the OBS websockets server.

Once the client is initialized, usage is as follows:
```
go client.Listen()
for event := range client.IncomingEvents {
	switch e := event.(type) {
	case *events.SomeEventA:
		...
	case *events.SomeEventB:
		...
	default:
	}
}
```
*/
func (c *Client) Listen() {
	var err error

	// separate connection
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

		switch knownEvent := events.GetEventForType(eventType); knownEvent {
		case nil:
			c.IncomingEvents <- unknownEvent
		default:
			if err := json.Unmarshal(raw, knownEvent); err != nil {
				log.Fatal(err)
			}
			c.IncomingEvents <- knownEvent
		}
	}
}

/*
Disconnect sends a message to the OBS websockets server to close any connections
we might have open. You don't really have to do this as any connections should
close when your program terminates or interrupts. But here's a function anyways.
*/
func (c *Client) Disconnect() error {
	f := func(conn *websocket.Conn) error {
		if conn != nil {
			return conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
			)
		}
		return nil
	}

	errors := []error{}
	for _, conn := range []*websocket.Conn{c.requestsConn, c.eventsConn} {
		if err := f(conn); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("Got the following errors disconnecting: %s", errors)
	}

	return nil
}
