package api

import (
	"encoding/json"
	"fmt"

	"github.com/andreykaipov/goobs/api/requests"
	"github.com/gorilla/websocket"
	uuid "github.com/nu7hatch/gouuid"
)

type Client struct {
	Conn *websocket.Conn
	//	IncomingEvents    chan events.Event
	//	IncomingResponses chan requests.Response
	IncomingEvents    chan json.RawMessage
	IncomingResponses chan json.RawMessage
}

// Conn is the backing websocket connection to the OBS websockets
// server. It's an exported field just for consistency with the Request
// subclients as they also have an exported `Conn` member.

// WriteMessage abstracts the logic every subclient uses to send a request and
// receive the corresponding response.
func (c *Client) WriteMessage(params requests.Params, response requests.Response) error {
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
	// Interestingly, it does seem thread-safe if I use a totally different
	// connection, in that connection A won't get a response from OBS for
	// a request from connection B. So, message IDs must be unique per
	// client? Events also appear to be broadcast to every client as they
	// have no associated message ID. Typing that out, it seems pretty
	// apparent, and likely the nature of WebSockets, and not just specific
	// to OBS. Who knows?
	//	for {
	//		if err := c.Conn.ReadJSON(response); err != nil {
	//			return err
	//		}
	//
	//		if response.GetMessageID() == params.GetMessageID() {
	//			break
	//		}
	//	}
	for msg := range c.IncomingResponses {
		if err := json.Unmarshal(msg, response); err != nil {
			return fmt.Errorf("Couldn't unmarshal message into an unknown event: %s", err)
		}
		//fmt.Printf("%#v\n", response)

		if response.GetMessageID() == params.GetMessageID() {
			break
		}
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
