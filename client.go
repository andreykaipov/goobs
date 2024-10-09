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
	"sync"
	"time"

	"github.com/andreykaipov/goobs/api"
	"github.com/andreykaipov/goobs/api/closecodes"
	"github.com/andreykaipov/goobs/api/events"
	"github.com/andreykaipov/goobs/api/events/subscriptions"
	"github.com/andreykaipov/goobs/api/opcodes"
	"github.com/gorilla/websocket"
	"github.com/hashicorp/logutils"
	"github.com/mmcloughlin/profile"
)

// Client represents a client to an OBS websockets server.
type Client struct {
	// See [Client.Listen] for more information.
	IncomingEvents chan any

	client *api.Client
	Categories

	conn               *websocket.Conn
	scheme             string
	host               string
	password           string
	dialer             *websocket.Dialer
	requestHeader      http.Header
	eventSubscriptions int
	profiler           *profile.Profile
	once               sync.Once
}

// Option represents a functional option of a Client.
type Option func(*Client)

// WithPassword sets the password of a client.
func WithPassword(x string) Option {
	return func(o *Client) { o.password = x }
}

// WithEventSubscriptions specifies the events we'd like to subscribe to via
// [Client.Listen]. The value is a bitmask of any of the subscription values
// specified in [subscriptions]. By default, all event categories are
// subscribed, except for events marked as high volume. High volume events must
// be explicitly subscribed to.
func WithEventSubscriptions(x int) Option {
	return func(o *Client) { o.eventSubscriptions = x }
}

// WithLogger sets the logger this library will use. See the logger.Logger
// interface. Should be compatible with most third-party loggers.
func WithLogger(x api.Logger) Option {
	return func(o *Client) { o.client.Log = x }
}

// WithDialer sets the underlying [gorilla/websocket.Dialer] should one want to
// customize things like the handshake timeout or TLS configuration. If this is
// not set, it'll use the provided [gorilla/websocket.DefaultDialer].
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
	return func(o *Client) { o.client.ResponseTimeout = time.Duration(x) }
}

// WithScheme sets the protocol scheme to use when connecting to the server,
// e.g. "ws" or "wss". The default is "ws". Please note however that the
// obs-websocket server does not currently support connecting over wss (ref:
// https://github.com/obsproject/obs-websocket/issues/26). To be able to
// connect over wss, you'll need to firest set up a reverse proxy in front of
// the server. The obs-websocket folks have a guide here:
// https://github.com/obsproject/obs-websocket/wiki/SSL-Tunneling.
func WithScheme(x string) Option {
	return func(o *Client) { o.scheme = x }
}

/*
Disconnect sends a message to the OBS websocket server to close the client's
open connection. You should defer a disconnection after creating your client to
ensure the program doesn't leave any lingering connections open and potentially
leak memory, e.g.:

	client, err := goobs.New("localhost:4455", goobs.WithPassword("secret"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect()
*/
func (c *Client) Disconnect() error {
	defer func() {
		if c.profiler != nil {
			c.client.Log.Printf("[DEBUG] Ending profiling")
			c.profiler.Stop()
		}
	}()

	c.client.Log.Printf("[DEBUG] Sending disconnect message")
	c.markDisconnected()

	if err := c.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Bye"),
	); err != nil {
		c.client.Log.Printf("[ERROR] Force closing connection: %s", err)
		return c.conn.Close()
	}

	return nil
}

func (c *Client) markDisconnected() {
	c.once.Do(func() {
		select {
		case c.client.Disconnected <- true:
		default:
		}

		c.client.Log.Printf("[TRACE] Closing internal channels")
		close(c.IncomingEvents)
		close(c.client.Opcodes)
		close(c.client.Disconnected)
	})
}

/*
New creates and configures a client to interact with the OBS websockets server.
It also opens up a connection, so be sure to check the error.
*/
func New(host string, opts ...Option) (*Client, error) {
	c := &Client{
		scheme:             "ws",
		host:               host,
		dialer:             websocket.DefaultDialer,
		requestHeader:      http.Header{"User-Agent": []string{"goobs/" + LibraryVersion}},
		eventSubscriptions: subscriptions.All,
		client: &api.Client{
			Disconnected:      make(chan bool),
			IncomingResponses: make(chan *opcodes.RequestResponse),
			Opcodes:           make(chan opcodes.Opcode),
			ResponseTimeout:   10000,
			Log: log.New(
				&logutils.LevelFilter{
					Levels:   []logutils.LogLevel{"TRACE", "DEBUG", "INFO", "ERROR", ""},
					MinLevel: logutils.LogLevel(strings.ToUpper(os.Getenv("GOOBS_LOG"))),
					Writer: api.LoggerWithWrite(func(p []byte) (int, error) {
						return os.Stderr.WriteString(fmt.Sprintf("\033[36m%s\033[0m", p))
					}),
				},
				"",
				log.Ltime|log.Lshortfile,
			),
		},
		IncomingEvents: make(chan any, 100),
	}

	if os.Getenv("GOOBS_PROFILE") != "" {
		c.profiler = profile.Start(
			profile.AllProfiles,
			profile.ConfigEnvVar("GOOBS_PROFILE"),
			profile.WithLogger(c.client.Log.(*log.Logger)),
		)
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
	u := url.URL{Scheme: c.scheme, Host: c.host}

	c.client.Log.Printf("[INFO] Connecting to %s", u.String())

	if c.conn, _, err = c.dialer.Dial(u.String(), c.requestHeader); err != nil {
		return err
	}

	authComplete := make(chan error)

	go c.handleRawServerMessages(authComplete)
	go func() {
		c.handleOpcodes(authComplete)
		close(c.client.IncomingResponses)
	}()

	timer := time.NewTimer(c.client.ResponseTimeout * time.Millisecond)
	defer timer.Stop()
	select {
	case a := <-authComplete:
		return a
	case <-timer.C:
		return fmt.Errorf("timeout waiting for authentication: %dms", c.client.ResponseTimeout)
	}
}

// translates raw server messages into opcodes
func (c *Client) handleRawServerMessages(auth chan<- error) {
	defer c.client.Log.Printf("[TRACE] Finished handling raw server messages")

	for {
		raw := json.RawMessage{}
		if err := c.conn.ReadJSON(&raw); err != nil {
			switch t := err.(type) {
			case *json.UnmarshalTypeError:
				c.client.Log.Printf("[ERROR] Reading from connection: %s: %s", t, raw)
				continue
			case *websocket.CloseError:
				switch t.Code {
				case closecodes.MessageDecodeError:
					continue
				case closecodes.MissingDataField:
					continue
				case closecodes.InvalidDataFieldType:
					continue
				case closecodes.InvalidDataFieldValue:
					continue
				case closecodes.UnknownOpCode:
					continue
				case closecodes.AuthenticationFailed:
					auth <- err
				}
				// this is like an "abrupt" disconnect
				c.client.Log.Printf("[DEBUG] Connection terminated: %s", t)
				c.markDisconnected()
				return
			default:
				switch t {
				case websocket.ErrCloseSent:
					// this is like a "graceful" disconnect (i think)
					c.client.Log.Printf("[DEBUG] Connection was closed: %s", t)
					c.markDisconnected()
				default:
					select {
					case <-c.client.Disconnected:
					default:
						c.client.Log.Printf("[ERROR] Unhandled error: %s", t)
					}
				}
				return
			}
		}

		c.client.Log.Printf("[TRACE] Raw server message: %s", raw)

		select {
		case <-c.client.Disconnected:
			// This might happen if the server sends messages to us
			// after we've already disconnected, e.g.:
			//
			// 1. client sends ToggleRecordPause request
			// 2. client gets the appropriate response for it
			// 3. client sends disconnect message immediately after
			// 4. client gets RecordStateChanged event
			c.client.Log.Printf("[DEBUG] Got %s from the server, but we've already disconnected!", raw)
			return
		default:
			opcode, err := opcodes.ParseRawMessage(raw)
			if err != nil {
				c.client.Log.Printf("[ERROR] Parse raw message: %s: %s", err, raw)
			}

			c.client.Opcodes <- opcode
		}
	}
}

// here's the meat of the operation
// handles both server and client opcodes
func (c *Client) handleOpcodes(auth chan<- error) {
	defer c.client.Log.Printf("[TRACE] Finished handling opcodes")

	for op := range c.client.Opcodes {
		switch val := op.(type) {

		case *opcodes.Hello:
			c.client.Log.Printf("[INFO] Hello; authenticating...")

			// we always try to auth; servers with auth
			// disabled will just ignore it if anything
			hash := sha256.Sum256([]byte(c.password + val.Authentication.Salt))
			secret := base64.StdEncoding.EncodeToString(hash[:])
			authHash := sha256.Sum256([]byte(secret + val.Authentication.Challenge))
			authSecret := base64.StdEncoding.EncodeToString(authHash[:])

			go func() {
				c.client.Opcodes <- &opcodes.Identify{
					RPCVersion:         val.RPCVersion,
					Authentication:     authSecret,
					EventSubscriptions: c.eventSubscriptions,
				}
			}()

		case *opcodes.Identify:
			c.client.Log.Printf("[INFO] Identify;")

			msg := opcodes.Wrap(val).Bytes()
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				auth <- fmt.Errorf("sending Identify to server `%s`: %w", msg, err)
			}

		case *opcodes.Identified:
			c.client.Log.Printf("[INFO] Identified; negotiated RPC version: %d", val.NegotiatedRPCVersion)
			auth <- nil

		case *opcodes.Reidentify:
			// can't imagine we need this

		case *opcodes.Event:
			c.client.Log.Printf("[TRACE] Got %s event: %s", val.Type, val.Data)

			event := events.GetType(val.Type)

			var data json.RawMessage
			if data = val.Data; data == nil {
				data = []byte("{}")
			}

			if err := json.Unmarshal(data, event); err != nil {
				c.client.Log.Printf("[ERROR] Unmarshalling `%s` into type %T: %s", val.Data, event, err)
			}

			c.writeEvent(event)

		case *opcodes.Request:
			c.client.Log.Printf("[TRACE] Got %s Request with ID %s", val.Type, val.ID)

			msg := opcodes.Wrap(val).Bytes()
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				c.client.Log.Printf("[ERROR] Sending Request to server `%s`: %s", msg, err)
			}

		case *opcodes.RequestResponse:
			c.client.Log.Printf("[TRACE] Got %s Response for ID %s (%d)", val.Type, val.ID, val.Status.Code)

			c.client.IncomingResponses <- val

		default:
			c.client.Log.Printf("[ERROR] Unhandled opcode %T", op)
		}
	}
}

// Since our events channel is buffered and might not necessarily be used, we
// purge old events and write latest ones so that whenever somebody might want
// to use it, they'll have the latest events available to them.
func (c *Client) writeEvent(event any) {
	select {
	case <-c.client.Disconnected:
	case c.IncomingEvents <- event:
	default:
		if len(c.IncomingEvents) == cap(c.IncomingEvents) {
			// incoming events was full (but might not be by now),
			// so safely read off the oldest, and write the latest
			select {
			case <-c.IncomingEvents:
			default:
			}

			c.IncomingEvents <- event
		}
	}
}

// Listen hooks into the incoming events from the server. You'll have to make
// type assertions to ensure you're getting the right events, e.g.:
//
//	client.Listen(func(event any) {
//		switch val := event.(type) {
//		case *events.InputVolumeChanged:
//			// event i was looking for
//		default:
//			// other events
//		}
//	})
//
// If looking for high volume events, make sure you've initialized the client
// with the appropriate subscriptions with [WithEventSubscriptions].
func (c *Client) Listen(f func(any)) {
	defer c.client.Log.Printf("[TRACE] Finished handling incoming events")

	for event := range c.IncomingEvents {
		f(event)
	}
}
