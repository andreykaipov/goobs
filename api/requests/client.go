package requests

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/andreykaipov/goobs/api/opcodes"
	"github.com/andreykaipov/goobs/apiv5/events"
	"github.com/buger/jsonparser"
	"github.com/gorilla/websocket"
	uuid "github.com/nu7hatch/gouuid"
)

// Logger is a interface compatible with both the stdlib's logger and some
// third-party loggers.
type Logger interface {
	Printf(string, ...interface{})
}

// Client represents a requests client to the OBS websocket server. Its
// intention is to provide a means of communication between the top-level client
// and the category-level clients, so while its fields are exported, they should
// be of no interest to consumers of this library.
type Client struct {
	// The backing websocket connection to the OBS websocket server.
	Conn *websocket.Conn

	// The time we're willing to wait to receive a response from the server.
	ResponseTimeout time.Duration

	// Events broadcast by the server when actions happen within OBS.
	IncomingEvents chan events.Event

	// Raw JSON message responses from the websocket server.
	IncomingResponses chan json.RawMessage

	Log Logger
}

/*
Disconnect sends a message to the OBS websocket server to close the client's
open connection. You don't really have to do this as any connections should
close when your program terminates or interrupts. But here's a function anyways.
*/
func (c *Client) Disconnect() error {
	close(c.IncomingResponses)
	close(c.IncomingEvents)

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
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	request := opcodes.Wrap(&opcodes.Request{
		RequestType: params.GetSelfName(),
		RequestID:   id.String(),
		RequestData: params,
	})
	if err := c.Conn.WriteMessage(websocket.TextMessage, request); err != nil {
		panic(fmt.Errorf("Couldn't write message: %w", err))
	}

	c.Log.Printf("Sent request: %s", request)

	var msg json.RawMessage
	select {
	case msg = <-c.IncomingResponses:
	case <-time.After(c.ResponseTimeout * time.Millisecond):
		return fmt.Errorf("Timed out receiving response from server for request %q", params.GetRequestType())
	}

	// TODO parse requestId and requestStatus as common response params

	c.Log.Printf("Raw response: %s", msg)

	responseData, _, _, err := jsonparser.Get(msg, "responseData")
	if err != nil {
		return fmt.Errorf("Couldn't find 'responseData' key from response: %s", err)
	}

	if err := json.Unmarshal(responseData, response); err != nil {
		return fmt.Errorf("Couldn't unmarshal message into an unknown event: %s", err)
	}

	//
	responseCode, err := jsonparser.GetInt(msg, "requestStatus", "code")
	if err != nil {
		return fmt.Errorf("Couldn't find 'requestStatus.code' key from response: %s", err)
	}
	responseComment, _ := jsonparser.GetString(msg, "requestStatus", "comment")
	responseMessageID, err := jsonparser.GetString(msg, "requestId")
	if err != nil {
		return fmt.Errorf("Couldn't find 'requestId' key from response: %s", err)
	}

	// request ID on response should be a mirror of what the client sent, so
	// this is good to check. however, I'm pretty sure a mismatch like this
	// could only happen in a concurrent context... and I'm pretty sure
	// gorilla/websocket would panic before then, so... 🤷
	if responseMessageID != id.String() {
		return fmt.Errorf(
			"Sent request %s, with message ID %s, but next response in channel has message ID %s",
			params.GetSelfName(),
			id,
			responseMessageID,
		)
	}

	if responseCode != 100 {
		return fmt.Errorf(
			"Got error from OBS executing request %q: %s",
			params.GetRequestType(),
			responseComment,
		)
	}

	return nil
}
