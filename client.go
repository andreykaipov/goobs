package goobs

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/andreykaipov/goobs/api"
	general "github.com/andreykaipov/goobs/api/requests/general"
	"github.com/gorilla/websocket"
)

// Client represents a client to an OBS websockets server.
type Client struct {
	api.Client

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
		host: host,
	}
	c.IncomingResponses = make(chan json.RawMessage, 100) //make(chan events.Event, 100),
	c.IncomingEvents = make(chan json.RawMessage, 100)    //make(chan events.Event, 100),

	for _, opt := range opts {
		opt(c)
	}

	c.Conn, err = c.connect()
	if err != nil {
		return nil, err
	}

	setClients(c)

	go c.Listen()

	if err := c.authenticate(); err != nil {
		return nil, err
	}

	return c, nil
}

// Pretty much the pseudo-code from
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#authentication
func (c *Client) authenticate() error {
	authReqResp, err := c.General.GetAuthRequired()
	if err != nil {
		return fmt.Errorf("Failed getting auth: %s", err)
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
		return fmt.Errorf("Failed authenticating: %s", err)
	}

	return nil
}

func (c *Client) Listen() {
	messages := make(chan json.RawMessage)
	errors := make(chan error)
	//go c.handleErrors(errors)
	go c.handleRawMessages(messages, errors)
	c.handleMessages(messages, errors)
}

// func (c *Client) handleErrors(errors chan error) {
// 	for err := range errors {
// 		c.IncomingEvents <- events.WrapError(err)
// 	}
// }

func (c *Client) handleRawMessages(messages chan json.RawMessage, errors chan error) {
	for {
		msg := json.RawMessage{}
		if err := c.Conn.ReadJSON(&msg); err != nil {
			errors <- fmt.Errorf("Couldn't read JSON from websocket connection: %s", err)
			continue
		}

		messages <- msg
	}
}

func (c *Client) handleMessages(messages chan json.RawMessage, errors chan error) {
	for raw := range messages {
		checked := map[string]interface{}{}
		if err := json.Unmarshal(raw, &checked); err != nil {
			errors <- fmt.Errorf("Couldn't unmarshal message: %s", err)
			continue
		}

		if _, ok := checked["update-type"]; ok {
			// Non-blocking write in case our events channel is
			// full, since user may not ever use it
			select {
			case c.IncomingEvents <- raw:
			default:
			}

			continue
		}

		if _, ok := checked["message-id"]; ok {
			c.IncomingResponses <- raw
			continue
		}

		panic("idk what kinda message this is lol")
		//		unknownEvent := &events.EventBasic{}
		//		if err := json.Unmarshal(raw, unknownEvent); err != nil {
		//			errors <- fmt.Errorf("Couldn't unmarshal message into an unknown event: %s", err)
		//			continue
		//		}
		//
		//		eventType := unknownEvent.UpdateType
		//
		//		switch knownEvent := events.GetEventForType(eventType); knownEvent {
		//		case nil:
		//			c.IncomingEvents <- unknownEvent
		//		default:
		//			if err := json.Unmarshal(raw, knownEvent); err != nil {
		//				errors <- fmt.Errorf("Couldn't unmarshal message into an event type of %q: %s", eventType, err)
		//				continue
		//			}
		//
		//			c.IncomingEvents <- knownEvent
		//		}
	}
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
/*
func (c *Client) Listen() {
	var err error

	c.eventingConn, err = c.connect() // separate connection
	if err != nil {
		panic(fmt.Errorf("Failed establishing an eventing connection: %s", err))
	}

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
*/

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
Disconnect sends a message to the OBS websockets server to close the client's
open connection. You don't really have to do this as any connections should
close when your program terminates or interrupts. But here's a function anyways.
*/
func (c *Client) Disconnect() error {
	return c.Conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	)
}
