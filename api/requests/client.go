package requests

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	uuid "github.com/nu7hatch/gouuid"
)

// Client represents a client to an OBS websockets server, used for requests.
type Client struct {
	// Conn is the backing websocket connection to the OBS websockets
	// server. It's exported to pass the underlying connection to each
	// category subclient. You shouldn't worry about this.
	Conn *websocket.Conn

	IncomingResponses chan json.RawMessage
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
//
// To get the response for a sent request, we can just read the next response
// from our channel. This works fine in a single-threaded context, and the
// message IDs of both the sent request and response should match.
//
// In a concurrent context, this isn't necessarily true, but since
// gorilla/websocket doesn't handle concurrency anyways, who cares? We could
// technically add a mutex in between sending our request and reading from the
// channel, but ehh...
//
// Interestingly, it does seem thread-safe if I use a totally different
// connection, in that connection A won't get a response from OBS for a request
// from connection B. So, message IDs must be unique per client? More
// interestingly, events appear to be broadcast to every client, maybe because
// they have no associated message ID?
func (c *Client) SendRequest(params Params, response Response) error {
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

	msg := <-c.IncomingResponses
	if err := json.Unmarshal(msg, response); err != nil {
		return fmt.Errorf("Couldn't unmarshal message into an unknown event: %s", err)
	}

	// gorilla/websocket would panic before this could happen in
	// a concurrent context, and idk how this could ever happen otherwise
	if response.GetMessageID() != params.GetMessageID() {
		return fmt.Errorf(
			"Sent request %s, with message ID %s, but next response in channel has message ID %s",
			params.Name(),
			params.GetMessageID(),
			response.GetMessageID(),
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
