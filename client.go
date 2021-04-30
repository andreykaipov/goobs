package goobs

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/andreykaipov/goobs/api/events"
	general "github.com/andreykaipov/goobs/api/requests/general"
	"github.com/gorilla/websocket"
)

// Client represents a client to an OBS websockets server.
type Client struct {
	IncomingEvents chan events.Event

	// We use two different connections for events and requests to avoid
	// race conditions when reading events and request responses
	eventingConn *websocket.Conn
	requestsConn *websocket.Conn

	host     string
	password string

	subclients
}

// Option represents a functional option of a Client.
type Option func(*Client)

// WithPassword sets the password of a client.
func WithPassword(x string) Option {
	return func(o *Client) {
		o.password = x
	}
}

/*
New creates and configures a client to interact with the OBS websockets server.
It also opens up a connection, so be sure to check the error.
*/
func New(host string, opts ...Option) (c *Client, err error) {
	c = &Client{
		IncomingEvents: make(chan events.Event, 100),
		host:           host,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.requestsConn, err = c.connect()
	if err != nil {
		return nil, err
	}

	setClients(c)

	if err := c.authenticate(); err != nil {
		return nil, fmt.Errorf("Failing authenticating: %s", err)
	}

	return c, nil
}

// Pretty much the pseudo-code from
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#authentication
func (c *Client) authenticate() error {
	authReqResp, err := c.General.GetAuthRequired()
	if err != nil {
		return err
	}

	if !authReqResp.AuthRequired {
		return nil
	}

	hash := sha256.Sum256([]byte(c.password + authReqResp.Salt))
	secret := base64.StdEncoding.EncodeToString(hash[:])

	authHash := sha256.Sum256([]byte(secret + authReqResp.Challenge))
	authSecret := base64.StdEncoding.EncodeToString(authHash[:])

	if _, err := c.General.Authenticate(&general.AuthenticateParams{
		Auth: authSecret,
	}); err != nil {
		return err
	}

	return nil
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
	c.eventingConn, err = c.connect()
	if err != nil {
		panic(fmt.Errorf("Failed establishing an eventing connection: %s", err))
	}

	// The eventing loop involves a few Go routines. We use this channel to
	// keep track of any errors so the client can find out about them, if it
	// wants to.
	messages := make(chan json.RawMessage)
	errors := make(chan error)
	go c.handleErrors(errors)
	go c.handleRawEvents(messages, errors)
	c.handleEvents(messages, errors)
}

func (c *Client) handleErrors(errors chan error) {
	for err := range errors {
		c.IncomingEvents <- events.WrapError(err)
	}
}

func (c *Client) handleRawEvents(messages chan json.RawMessage, errors chan error) {
	for {
		raw := json.RawMessage{}
		if err := c.eventingConn.ReadJSON(&raw); err != nil {
			errors <- fmt.Errorf("Couldn't read JSON from websocket connection: %s", err)
			continue
		}

		messages <- raw
	}
}

func (c *Client) handleEvents(messages chan json.RawMessage, errors chan error) {
	for raw := range messages {
		unknownEvent := &events.EventBasic{}
		if err := json.Unmarshal(raw, unknownEvent); err != nil {
			errors <- fmt.Errorf("Couldn't unmarshal message into an unknown event: %s", err)
			continue
		}

		eventType := unknownEvent.UpdateType

		switch knownEvent := events.GetEventForType(eventType); knownEvent {
		case nil:
			c.IncomingEvents <- unknownEvent
		default:
			if err := json.Unmarshal(raw, knownEvent); err != nil {
				errors <- fmt.Errorf("Couldn't unmarshal message into an event type of %q: %s", eventType, err)
				continue
			}

			c.IncomingEvents <- knownEvent
		}
	}
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
Disconnect sends a message to the OBS websockets server to close any connections
we might have open. You don't really have to do this as any connections should
close when your program terminates or interrupts. But here's a function anyways.
*/
func (c *Client) Disconnect() error {
	disconnect := func(conn *websocket.Conn) error {
		if conn != nil {
			return conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
			)
		}
		return nil
	}

	errors := []error{}
	for _, conn := range []*websocket.Conn{c.requestsConn, c.eventingConn} {
		if err := disconnect(conn); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("Got the following errors disconnecting: %s", errors)
	}

	return nil
}
