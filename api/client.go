package api

import (
	"encoding/json"
	"fmt"

	"github.com/andreykaipov/goobs/api/events"
	"github.com/andreykaipov/goobs/api/requests"
	"github.com/gorilla/websocket"
	uuid "github.com/nu7hatch/gouuid"
)

// Client represents a client to an OBS websockets server.
type Client struct {
	// Conn is the backing websocket connection to the OBS websockets
	// server. It's exported to pass the underlying connection to each
	// category subclient. You shouldn't worry about this.
	Conn *websocket.Conn

	// IncomingEvents is used to read events from OBS. For example,
	//
	// ```go
	// for event := range client.IncomingEvents {
	// 	switch e := event.(type) {
	// 	case *events.SomeEventA:
	// 		...
	// 	case *events.SomeEventB:
	// 		...
	// 	default:
	// 	}
	// }
	// ```
	IncomingEvents    chan events.Event
	IncomingResponses chan json.RawMessage
}

func New() Client {
	return Client{
		Conn:              nil,
		IncomingEvents:    make(chan events.Event),
		IncomingResponses: make(chan json.RawMessage),
	}
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

// SendRequest abstracts the logic every subclient uses to send a request and
// receive the corresponding response.
func (c *Client) SendRequest(params requests.Params, response requests.Response) error {
	params.SetRequestType(params.Name())

	if params.GetRequestType() == "" {
		return fmt.Errorf("Params %T has an empty RequestType", params)
	}

	if params.GetMessageID() == "" {
		id, err := uuid.NewV4()
		if err != nil {
			return err
		}

		params.SetMessageID(id.String())
	}

	if err := c.Conn.WriteJSON(params); err != nil {
		return err
	}

	// Some requests trigger an event, which we don't really care about, so
	// keep reading until we get a message with a matching message ID as
	// that'd be our response. Clearly, this isn't safe against concurrency,
	// since one goroutine might accidentally read a response belonging to
	// a request from some other goroutine, which would cause that goroutine
	// to never receive a response. However, gorilla/websocket doesn't
	// handle concurrency either, so who cares?
	//

	// Read the next response from our incoming responses. In
	// a single-threaded context, the message IDs of both the sent request
	// and the response should always match.
	//
	// In a concurrent context, this isn't necessarily true, but since
	// gorilla/websocket doesn't handle concurrency anyways, who cares?
	// We could technically add a mutex in between sending our request and
	// reading from this channel, but ehh...
	//
	// Interestingly, it does seem thread-safe if I use a totally different
	// connection, in that connection A won't get a response from OBS for
	// a request from connection B. So, message IDs must be unique per
	// client? However, events appear to be broadcast to every client as
	// they have no associated message ID. Typing that out, it seems pretty
	// apparent, and likely the nature of WebSockets, and not just specific
	// to OBS. Who knows?

	msg := <-c.IncomingResponses
	if err := json.Unmarshal(msg, response); err != nil {
		return fmt.Errorf("Couldn't unmarshal message into an unknown event: %s", err)
	}

	if response.GetMessageID() != params.GetMessageID() {
		return fmt.Errorf(
			"Sent request %s, with message ID %s, but next response in channel has message ID %s",
			params.Name(), params.GetMessageID(), response.GetMessageID(),
		)
	}

	if response.GetStatus() == "error" {
		return fmt.Errorf(
			"Got error from OBS executing request %q: %s",
			params.GetRequestType(),
			response.GetError(),
		)
	}

	return nil
}
