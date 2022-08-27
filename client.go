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
	"strings"
	"time"

	"github.com/andreykaipov/goobs/api"
	"github.com/andreykaipov/goobs/api/events"
	"github.com/andreykaipov/goobs/api/events/subscriptions"
	"github.com/andreykaipov/goobs/api/opcodes"
	"github.com/buger/jsonparser"
	"github.com/gorilla/websocket"
	"github.com/hashicorp/logutils"
)

// Client represents a client to an OBS websockets server.
type Client struct {
	*api.Client
	subclients
	conn               *websocket.Conn
	host               string
	password           string
	dialer             *websocket.Dialer
	requestHeader      http.Header
	eventSubscriptions int
	errors             chan error
}

// Option represents a functional option of a Client.
type Option func(*Client)

// WithPassword sets the password of a client.
func WithPassword(x string) Option {
	return func(o *Client) { o.password = x }
}

// WithEventSubscriptions specifies the events we'd like to susbcribe to via
// `client.Listen()`. The value is a bitmask of any of the subscription values
// specified in api/events/subscriptions. By default, all event categories are
// subscribed, except for events marked as high volume. High volume events must
// be explicitly subscribed to.
func WithEventSubscriptions(x int) Option {
	return func(o *Client) { o.eventSubscriptions = x }
}

// WithLogger sets the logger this library will use. See the logger.Logger
// interface. Should be compatible with most third-party loggers.
func WithLogger(x api.Logger) Option {
	return func(o *Client) { o.Log = x }
}

// WithDialer sets the underlying Gorilla WebSocket Dialer (see
// https://pkg.go.dev/github.com/gorilla/websocket#Dialer), should one want to
// customize things like the handshake timeout or TLS configuration. If this is
// not set, it'll use the provided DefaultDialer (see
// https://pkg.go.dev/github.com/gorilla/websocket#pkg-variables).
func WithDialer(x *websocket.Dialer) Option {
	return func(o *Client) { o.dialer = x }
}

// WithRequestHeader sets custom headers our client can send when trying to
// connect to the WebSockets server, allowing us specify the origin,
// subprotocols, or the user agent.
func WithRequestHeader(x http.Header) Option {
	return func(o *Client) { o.requestHeader = x }
}

// WithResponseTimeout sets the time we're willing to wait to receive a response
// from the server for any request, before responding with an error. It's in
// milliseconds. The default timeout is 10 seconds.
func WithResponseTimeout(x time.Duration) Option {
	return func(o *Client) { o.ResponseTimeout = time.Duration(x) }
}

/*
Disconnect sends a message to the OBS websocket server to close the client's
open connection. You don't really have to do this as any connections should
close when your program terminates or interrupts. But here's a function anyways.
*/
func (c *Client) Disconnect() error {
	c.Log.Printf("[DEBUG] Sending disconnect message")

	return c.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Bye"),
	)
}

/*
New creates and configures a client to interact with the OBS websockets server.
It also opens up a connection, so be sure to check the error.
*/
func New(host string, opts ...Option) (*Client, error) {
	c := &Client{
		host:               host,
		dialer:             websocket.DefaultDialer,
		requestHeader:      http.Header{"User-Agent": []string{"goobs/" + goobs_version}},
		eventSubscriptions: subscriptions.All,
		errors:             make(chan error),
		Client: &api.Client{
			IncomingEvents:    make(chan interface{}, 100),
			IncomingResponses: make(chan *opcodes.RequestResponse),
			Opcodes:           make(chan opcodes.Opcode),
			ResponseTimeout:   10000,
			Log: log.New(
				&logutils.LevelFilter{
					Levels:   []logutils.LogLevel{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", ""},
					MinLevel: logutils.LogLevel(strings.ToUpper(os.Getenv("GOOBS_LOG"))),
					Writer: api.LoggerWithWrite(func(p []byte) (int, error) {
						return os.Stderr.WriteString(fmt.Sprintf("\033[36m%s\033[0m", p))
					}),
				},
				"",
				log.Ltime|log.Lshortfile,
			),
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	if err := c.connect(); err != nil {
		return nil, err
	}

	setClients(c)

	return c, nil
}

func (c *Client) connect() (err error) {
	if err = c.checkProtocolVersion(); err != nil {
		return
	}

	u := url.URL{Scheme: "ws", Host: c.host}

	c.Log.Printf("[INFO] Connecting to %s", u.String())

	if c.conn, _, err = c.dialer.Dial(u.String(), c.requestHeader); err != nil {
		return err
	}

	authComplete := make(chan error)

	go c.handleErrors()
	go c.handleRawServerMessages(authComplete)
	go c.handleOpcodes(authComplete)

	select {
	case a := <-authComplete:
		return a
	case <-time.After(c.ResponseTimeout * time.Millisecond):
		return fmt.Errorf("timeout waiting for authentication: %dms", c.ResponseTimeout)
	}
}

// Here we connect and send a v4.x-like message to the server. If it responds
// properly, we want to fail immediately. Motiviation is similar to what the
// obs-websocket server does when it detects a v4 client connecting to it (see
// ref). However, in this case we're checking a v5 client against a v4 server.
// We use a separate connection to not error out the main one, because a v5
// server will be disappointed with a v4 message.
//
// ref:
// https://github.com/obsproject/obs-websocket/blob/5.0.0/src/websocketserver/WebSocketServer.cpp#L431-L439
func (c *Client) checkProtocolVersion() error {
	c.Log.Printf("[DEBUG] Checking correct protocol version")

	u := url.URL{Scheme: "ws", Host: c.host}
	conn, _, err := c.dialer.Dial(u.String(), c.requestHeader)
	if err != nil {
		return err
	}
	defer conn.Close()

	_ = conn.WriteMessage(
		websocket.TextMessage,
		[]byte(`{"request-type":"GetAuthRequired","message-id":"1"}`),
	)

	raw := json.RawMessage{}
	_ = conn.ReadJSON(&raw)
	c.Log.Printf("[DEBUG] %s", raw)

	var val []byte
	if val, _, _, _ = jsonparser.Get(raw, "authRequired"); val != nil {
		return fmt.Errorf(
			"%s; %s",
			"obs-websocket v4.x is not supported",
			"please use github.com/andreykaipov/goobs@v0.8.1",
		)
	}

	return nil
}

// expose errors as events 🤷
func (c *Client) handleErrors() {
	for err := range c.errors {
		c.Log.Printf("[ERROR] %s", err)
		c.writeEvent(err)
	}
}

// translates raw server messages into opcodes
func (c *Client) handleRawServerMessages(auth chan<- error) {
	for {
		raw := json.RawMessage{}
		if err := c.conn.ReadJSON(&raw); err != nil {
			switch t := err.(type) {
			case *websocket.CloseError:
				// for values the close error might have, see
				// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#websocketclosecode
				switch t.Code {
				case websocket.CloseNormalClosure:
					c.Log.Printf("[DEBUG] Closing connection: %s", t.Text)
				case 4009: // WebSocketCloseCode::AuthenticationFailed
					c.Log.Printf("[INFO] Closing connection: %s", t.Text)
					auth <- err
				default:
					c.errors <- fmt.Errorf("reading raw message: closed: %w", t)
				}
				return
			default:
				c.errors <- fmt.Errorf("reading raw message from websocket connection: %w", t)
				continue
			}
		}

		c.Log.Printf("[DEBUG] Raw server message: %s", raw)

		opcode, err := opcodes.ParseRawMessage(raw)
		if err != nil {
			c.errors <- fmt.Errorf("parse raw message: %w", err)
		}

		c.Opcodes <- opcode
	}
}

// here's the meat of the operation
// handles both server and client opcodes
func (c *Client) handleOpcodes(auth chan<- error) {
	for op := range c.Opcodes {
		switch val := op.(type) {

		case *opcodes.Hello:
			c.Log.Printf("[INFO] Hello; authenticating...")

			// we always try to auth; servers with auth
			// disabled will just ignore it if anything
			hash := sha256.Sum256([]byte(c.password + val.Authentication.Salt))
			secret := base64.StdEncoding.EncodeToString(hash[:])
			authHash := sha256.Sum256([]byte(secret + val.Authentication.Challenge))
			authSecret := base64.StdEncoding.EncodeToString(authHash[:])

			go func() {
				c.Opcodes <- &opcodes.Identify{
					RPCVersion:         val.RPCVersion,
					Authentication:     authSecret,
					EventSubscriptions: c.eventSubscriptions,
				}
			}()

		case *opcodes.Identify:
			c.Log.Printf("[INFO] Identify;")

			msg := opcodes.Wrap(val).Bytes()
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				auth <- fmt.Errorf("sending Identify to server `%s`: %w", msg, err)
			}

		case *opcodes.Identified:
			c.Log.Printf("[INFO] Identified; negotiated RPC version: %d", val.NegotiatedRPCVersion)
			auth <- nil

		case *opcodes.Reidentify:
			// can't imagine we need this

		case *opcodes.Event:
			c.Log.Printf("[INFO] Got %s Event", val.Type)
			c.Log.Printf("[DEBUG] Event Data: %s", val.Data)

			event := events.GetType(val.Type)

			var data json.RawMessage
			if data = val.Data; data == nil {
				data = []byte("{}")
			}

			if err := json.Unmarshal(data, event); err != nil {
				c.errors <- fmt.Errorf(
					"unmarshalling `%s` into type %T: %s",
					val.Data,
					event,
					err,
				)
			}

			c.writeEvent(event)

		case *opcodes.Request:
			c.Log.Printf("[DEBUG] Got %s Request with ID %s", val.Type, val.ID)

			msg := opcodes.Wrap(val).Bytes()
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				c.errors <- fmt.Errorf("sending Request to server `%s`: %w", msg, err)
			}

		case *opcodes.RequestResponse:
			c.Log.Printf("[INFO] Got %s Response for ID %s (%d)", val.Type, val.ID, val.Status.Code)

			c.IncomingResponses <- val

		default:
			c.errors <- fmt.Errorf("unhandled opcode %T", op)
		}
	}
}

// Since our events channel is buffered and might not necessarily be used, we
// purge old events and write latest ones so that whenever somebody might want
// to use it, they'll have the latest events available to them.
func (c *Client) writeEvent(event interface{}) {
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

func (c *Client) Listen(f func(interface{})) {
	for event := range c.IncomingEvents {
		f(event)
	}
}
