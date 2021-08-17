package goobs

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/andreykaipov/goobs/api/events"
	"github.com/andreykaipov/goobs/api/requests"
	general "github.com/andreykaipov/goobs/api/requests/general"
	"github.com/gorilla/websocket"
)

// Client represents a client to an OBS websockets server.
type Client struct {
	IncomingEvents chan events.Event
	*requests.Client

	subclients
	host     string
	password string
	debug    *bool
}

// Option represents a functional option of a Client.
type Option func(*Client)

// WithPassword sets the password of a client.
func WithPassword(x string) Option {
	return func(o *Client) {
		o.password = x
	}
}

// WithDebug enables debug logging via a default logger.
func WithDebug(x bool) Option {
	return func(o *Client) {
		o.debug = &x
	}
}

// WithLogger sets the logger to use for debug logging. Providing a logger
// implicitly turns debug logging on, unless debug logging is explicitly
// disabled.
func WithLogger(x requests.Logger) Option {
	return func(o *Client) {
		o.Log = x
	}
}

type discard struct{}

func (o *discard) Printf(format string, v ...interface{}) {}

type coloredStderr struct{}

func (o *coloredStderr) Write(p []byte) (n int, err error) {
	return os.Stderr.WriteString(fmt.Sprintf("\033[36m%s\033[0m", p))
}

/*
New creates and configures a client to interact with the OBS websockets server.
It also opens up a connection, so be sure to check the error.
*/
func New(host string, opts ...Option) (*Client, error) {
	c := &Client{
		Client: &requests.Client{},
		host:   host,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.Log == nil && c.debug == nil {
		c.Log = &discard{}
	}
	if c.Log == nil && *c.debug {
		c.Log = log.New(&coloredStderr{}, "goobs/debug: ", log.Lshortfile|log.LstdFlags)
	}
	if c.debug != nil && !*c.debug {
		c.Log = &discard{}
	}

	if err := c.connect(); err != nil {
		return nil, err
	}

	setClients(c)

	c.IncomingEvents = make(chan events.Event, 100)
	c.IncomingResponses = make(chan json.RawMessage, 100)
	go c.handleMessages()

	if err := c.authenticate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) connect() (err error) {
	u := url.URL{Scheme: "ws", Host: c.host}

	c.Log.Printf("Connecting to %s", u.String())

	if c.Conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil); err != nil {
		return err
	}

	return nil
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

func (c *Client) handleMessages() {
	messages := make(chan json.RawMessage)
	errors := make(chan error)
	go c.handleErrors(errors)
	go c.handleConnection(messages, errors)
	c.handleRawMessages(messages, errors)
}

// Expose eventing errors as... more events
func (c *Client) handleErrors(errors chan error) {
	for err := range errors {
		c.IncomingEvents <- events.WrapError(err)
	}
}

func (c *Client) handleConnection(messages chan json.RawMessage, errors chan error) {
	for {
		msg := json.RawMessage{}
		if err := c.Conn.ReadJSON(&msg); err != nil {
			errors <- fmt.Errorf("Couldn't read JSON from websocket connection: %s", err)
			continue
		}

		messages <- msg
	}
}

func (c *Client) handleRawMessages(messages chan json.RawMessage, errors chan error) {
	for raw := range messages {
		// Parse into a generic map to figure out if it's an event or
		// a response to a request first. Then act accordingly.
		checked := map[string]interface{}{}
		if err := json.Unmarshal(raw, &checked); err != nil {
			errors <- fmt.Errorf("Couldn't unmarshal message: %s", err)
			continue
		}

		// Responses are parsed when the request is sent
		if _, ok := checked["message-id"]; ok {
			c.IncomingResponses <- raw
			continue
		}

		// Events are parsed immediately. Kinda wasteful to do since
		// they might not ever be read, but it's not the end of the
		// world. Could always add an explicit option to enable events!
		if _, ok := checked["update-type"]; ok {
			event, err := events.Parse(raw)
			if err != nil {
				errors <- fmt.Errorf("Couldn't parse raw event: %s", err)
				continue
			}

			select {
			case c.IncomingEvents <- event:
			default:
			}

			continue
		}

		panic("idk what kinda message this is lol")
	}
}
