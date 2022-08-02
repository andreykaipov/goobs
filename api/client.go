package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/andreykaipov/goobs/api/opcodes"
	"github.com/andreykaipov/goobs/api/requests"
	uuid "github.com/nu7hatch/gouuid"
)

type ResponsePair struct {
	*opcodes.RequestResponse
	ResponseType interface{}
}

// Client represents a requests client to the OBS websocket server. Its
// intention is to provide a means of communication between the top-level client
// and the category-level clients, so while its fields are exported, they should
// be of no interest to consumers of this library.
type Client struct {
	// The time we're willing to wait to receive a response from the server.
	ResponseTimeout time.Duration

	IncomingEvents    chan interface{}
	IncomingResponses chan *ResponsePair
	Opcodes           chan opcodes.Opcode
	Log               Logger
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
//
//
// https://github.com/obsproject/obs-websocket/blob/master/docs/generated/protocol.md#request-opcode-6
//
func (c *Client) SendRequest(name string, params interface{}) (interface{}, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	id := uid.String()

	c.Log.Printf("[INFO] Sending %s Request with ID %s", name, id)

	c.Opcodes <- &opcodes.Request{
		Type: name,
		ID:   id,
		Data: params,
	}

	var pair *ResponsePair
	select {
	case pair = <-c.IncomingResponses:
	case <-time.After(c.ResponseTimeout * time.Millisecond):
		return nil, fmt.Errorf("request %s: timeout waiting for response from server", name)
	}

	response := pair.RequestResponse
	responseType := pair.ResponseType

	// request ID on response should be a mirror of what the client sent, so
	// this is good to check. however, I'm pretty sure a mismatch like this
	// could only happen in a concurrent context... and I'm pretty sure
	// gorilla/websocket would panic before then, so... 🤷
	if response.ID != id {
		return nil, fmt.Errorf(
			"request %s: mismatched ID: expected response with ID %s, but got %s",
			name,
			id,
			response.ID,
		)
	}

	status := response.Status

	if code := status.Code; code != 100 {
		return nil, fmt.Errorf(
			"request %s: %s (%d): %s",
			name,
			requests.GetStatusForCode(code),
			code,
			status.Comment,
		)
	}

	data := response.Data

	if err := json.Unmarshal(data, responseType); err != nil {
		return nil, fmt.Errorf(
			"request %s: unmarshalling `%s` into type %T: %s",
			name,
			data,
			responseType,
			err,
		)
	}

	return responseType, nil
}
