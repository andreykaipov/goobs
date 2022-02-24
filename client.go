package goobs

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/andreykaipov/goobs/api/events"
	"github.com/andreykaipov/goobs/api/opcodes"
	"github.com/andreykaipov/goobs/api/requests"
	general "github.com/andreykaipov/goobs/api/requests/general"
	"github.com/gorilla/websocket"
)

var version = "0.9.0-dev"

// Client represents a client to an OBS websockets server.
type Client struct {
	*requests.Client
	subclients
	host          string
	password      string
	debug         *bool
	dialer        *websocket.Dialer
	requestHeader http.Header
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

// WithDialer sets the underlying Gorilla WebSocket Dialer (see
// https://pkg.go.dev/github.com/gorilla/websocket#Dialer), should one want to
// customize things like the handshake timeout or TLS configuration. If this is
// not set, it'll use the provided DefaultDialer (see
// https://pkg.go.dev/github.com/gorilla/websocket#pkg-variables).
func WithDialer(x *websocket.Dialer) Option {
	return func(o *Client) {
		o.dialer = x
	}
}

// WithRequestHeader sets custom headers our client can send when trying to
// connect to the WebSockets server, allowing us specify the origin,
// subprotocols, or the user agent.
func WithRequestHeader(x http.Header) Option {
	return func(o *Client) {
		o.requestHeader = x
	}
}

// WithResponseTimeout sets the time we're willing to wait to receive a response
// from the server for any request, before responding with an error. It's in
// milliseconds. The default timeout is 10 seconds.
func WithResponseTimeout(x time.Duration) Option {
	return func(o *Client) {
		o.ResponseTimeout = time.Duration(x)
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
		Client: &requests.Client{
			ResponseTimeout: 10000,
		},
		host: host,
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

	if c.dialer == nil {
		c.dialer = websocket.DefaultDialer
	}
	if c.requestHeader == nil {
		c.requestHeader = http.Header{
			"User-Agent": []string{"goobs/" + version},
		}
	}

	if err := c.connect(); err != nil {
		return nil, err
	}

	// setClients(c)

	// c.IncomingEvents = make(chan events.Event, 100)
	// c.IncomingResponses = make(chan json.RawMessage, 100)
	// go c.handleMessages()
	//
	//	if err := c.wrappedAuthentication(); err != nil {
	//		return nil, fmt.Errorf("Failed auth: %s", err)
	//	}

	return c, nil
}

func (c *Client) read() (json.RawMessage, *opcodes.Message, error) {
	unchecked := json.RawMessage{}

	if err := c.Conn.ReadJSON(&unchecked); err != nil {
		switch err.(type) {
		case *websocket.CloseError:
			return nil, nil, fmt.Errorf("Websocket connection closed: %s", err)
		default:
			return nil, nil, fmt.Errorf("Couldn't read JSON from websocket connection: %s", err)
		}
	}

	checked := &opcodes.Message{}

	if err := json.Unmarshal(unchecked, &checked); err != nil {
		return nil, nil, fmt.Errorf("Couldn't unmarshal message: %s", err)
	}

	return unchecked, checked, nil
}

func (c *Client) connect() (err error) {
	u := url.URL{Scheme: "ws", Host: c.host}

	c.Log.Printf("Connecting to %s", u.String())

	if c.Conn, _, err = c.dialer.Dial(u.String(), c.requestHeader); err != nil {
		return err
	}

	go func() {
		for {
			raw, msg, err := c.read()
			if err != nil {
				panic(err)
			}

			fmt.Println(string(raw))

			// process whatever opcodes we might get from the server
			//
			switch msg.Op {
			case 0:
				opcode := &opcodes.Hello{}
				if err := json.Unmarshal(msg.D, opcode); err != nil {
					panic(fmt.Errorf("Couldn't unmarshal message: %s", err))
				}

				c.Log.Printf("Got Hello; authenticating...")

				// we always try to auth; servers with auth
				// disabled will just ignore it if anything
				hash := sha256.Sum256([]byte(c.password + opcode.Authentication.Salt))
				secret := base64.StdEncoding.EncodeToString(hash[:])
				authHash := sha256.Sum256([]byte(secret + opcode.Authentication.Challenge))
				authSecret := base64.StdEncoding.EncodeToString(authHash[:])

				identify := opcodes.Wrap(&opcodes.Identify{
					RPCVersion:     opcode.RPCVersion,
					Authentication: authSecret,
				})
				if err := c.Conn.WriteMessage(websocket.TextMessage, identify); err != nil {
					panic(err)
				}
			case 2:
				opcode := &opcodes.Identified{}
				if err := json.Unmarshal(msg.D, opcode); err != nil {
					panic(fmt.Errorf("Couldn't unmarshal message: %s", err))
				}

				c.Log.Printf("Now connected; negotiated RPC version: %d", opcode.NegotiatedRPCVersion)
			case 5:
				// event
			case 7:
				// response to a request
			case 9:
				// response to a batch of requests
				// no clue what this is
			default:
				panic(fmt.Errorf("unhandled opcode %d", msg.Op))
			}
		}
	}()

	// Parse into a generic map to figure out the opcode.
	// Then act accordingly.
	//	if err := json.Unmarshal(msg, &checked); err != nil {
	//		panic(fmt.Errorf("Couldn't unmarshal message: %s", err))
	//	}

	//fmt.Printf("%#v\n", checked)

	//	// Responses are parsed in the embedded Client's `SendRequest`
	//	if _, ok := checked["message-id"]; ok {
	//		c.IncomingResponses <- raw
	//		continue
	//	}

	//	messages <- msg
	//	}()

	return nil
}

// Handling authentication errors is a tad tricky. Because the auth request we
// send depends on the eventing loop too, we need a way to return any errors
// that might come up when parsing the auth response, while also handling
// expected auth errors like bad creds.
func (c *Client) wrappedAuthentication() error {
	go func() {
		if err := c.authenticate(); err != nil {
			c.IncomingEvents <- events.WrapError(err)
		}
		c.IncomingEvents <- nil
	}()

	switch e := (<-c.IncomingEvents).(type) {
	case *events.Error:
		// this error can be from the above `authenticate()`, or from
		// any errors that might've come up during the eventing loop
		return e.Err
	case nil:
		return nil
	default:
		// only events as of now should be errors or our above nil
		return fmt.Errorf("Surely impossible? How did the server send actual events before authentication?")
	}
}

// Pretty much the pseudo-code from
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#authentication
func (c *Client) authenticate() error {
	authReqResp, err := c.General.GetAuthRequired()
	if err != nil {
		return fmt.Errorf("Failed getting auth required: %s", err)
	}

	if !authReqResp.AuthRequired {
		return nil
	}

	hash := sha256.Sum256([]byte(c.password + authReqResp.Salt))
	secret := base64.StdEncoding.EncodeToString(hash[:])

	authHash := sha256.Sum256([]byte(secret + authReqResp.Challenge))
	authSecret := base64.StdEncoding.EncodeToString(authHash[:])

	_, err = c.General.Authenticate(&general.AuthenticateParams{Auth: authSecret})

	return err
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
		c.writeEvent(events.WrapError(err))
	}
}

func (c *Client) handleConnection(messages chan json.RawMessage, errors chan error) {
	for {
		msg := json.RawMessage{}
		if err := c.Conn.ReadJSON(&msg); err != nil {
			switch err.(type) {
			case *websocket.CloseError:
				errors <- fmt.Errorf("Websocket connection closed: %s", err)
				return
			default:
				errors <- fmt.Errorf("Couldn't read JSON from websocket connection: %s", err)
				continue
			}
		}

		messages <- msg
	}
}

// Handles messages from the server. They might be response bodies associated
// with requests, or they can be events we can subscribe to via the
// `client.IncomingEvents` channel. Or they can be something totally else, in
// which case we expose the errors as more events! Despite also handling
// incoming responses, we refer to this loop as the "eventing loop" elsewhere in
// the comments.
func (c *Client) handleRawMessages(messages chan json.RawMessage, errors chan error) {
	for raw := range messages {
		// Parse into a generic map to figure out if it's an event or
		// a response to a request first. Then act accordingly.
		checked := map[string]interface{}{}
		if err := json.Unmarshal(raw, &checked); err != nil {
			errors <- fmt.Errorf("Couldn't unmarshal message: %s", err)
			continue
		}

		// Responses are parsed in the embedded Client's `SendRequest`
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

			c.writeEvent(event)
			continue
		}

		errors <- fmt.Errorf("Client/server version mismatch? Unrecognized message: %s", raw)
	}
}

// Since our events channel is buffered and might not necessarily be used, we
// purge old events and write latest ones so that whenever somebody might want
// to use it, they'll have the latest events available to them.
func (c *Client) writeEvent(event events.Event) {
	select {
	case c.IncomingEvents <- event:
	default:
		if len(c.IncomingEvents) == cap(c.IncomingEvents) {
			// incoming events was full (but might not be by now),
			// so safely read off the oldest, and write the latest
			select {
			case _ = <-c.IncomingEvents:
			default:
			}

			c.IncomingEvents <- event
		}
	}
}
