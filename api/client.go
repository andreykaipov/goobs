package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/andreykaipov/goobs/api/opcodes"
	"github.com/andreykaipov/goobs/api/requests"
	uuid "github.com/nu7hatch/gouuid"
)

type Params interface{ GetRequestName() string }

// Client represents a requests client to the OBS websocket server. Its
// intention is to provide a means of communication between the top-level client
// and the category-level clients, so while its fields are exported, they should
// be of no interest to consumers of this library.
type Client struct {
	// The time we're willing to wait to receive a response from the server.
	ResponseTimeout time.Duration

	IncomingEvents    chan interface{}
	IncomingResponses chan *opcodes.RequestResponse
	Opcodes           chan opcodes.Opcode
	Log               Logger
}

// SendRequest abstracts the logic every subclient uses to send a request and
// receive the corresponding response.
//
// To get the response for a sent request, we can just read the next response
// from our channel. This works fine in a single-threaded context, and the
// message IDs of both the sent request and response should match. In
// a concurrent context, this isn't necessarily true, but since
// gorilla/websocket doesn't handle concurrency (it'll panic; see
// https://github.com/gorilla/websocket/issues/119), who cares?
//
// Technically a request ID and response ID mismatch could happen if the server
// processes requests in a different order it received them (e.g. we should 1,
// then 2; but it processes 2, and then 1), then yeah... there'll be an error.
// We could add a mutex wrapping sending our request and reading from the
// channel, but I personally haven't experienced this yet, so
//
// It should be noted multiple connections to the server are totally fine.
// Phrased differently, mesasge IDs are unique per client. Moreover, events will
// be broadcast to every client.
func (c *Client) SendRequest(requestBody Params, responseBody interface{}) error {
	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	name := requestBody.GetRequestName()
	id := uid.String()

	c.Log.Printf("[INFO] Sending %s Request with ID %s", name, id)

	c.Opcodes <- &opcodes.Request{
		Type: name,
		ID:   id,
		Data: requestBody,
	}

	var response *opcodes.RequestResponse
	select {
	case response = <-c.IncomingResponses:
	case <-time.After(c.ResponseTimeout * time.Millisecond):
		return fmt.Errorf("request %s: timeout waiting for response from server", name)
	}

	// i'm being overly cautious here making sure the request ID on the
	// response mirrors what the client sent in the request. see the
	// function header comment regarding concurrency concerns.
	if response.ID != id {
		return fmt.Errorf(
			"request %s: mismatched ID: expected response with ID %s, but got %s",
			name,
			id,
			response.ID,
		)
	}

	status := response.Status

	if code := status.Code; code != 100 {
		return fmt.Errorf(
			"request %s: %s (%d)%s",
			name,
			requests.GetStatusForCode(code),
			code,
			func() (s string) {
				if status.Comment != "" {
					s = ": " + status.Comment
				}
				return
			}(),
		)
	}

	// some requests don't have any response fields, and if so they will
	// return nothing, not even `{}`, so we add that ourselves so the
	// unmarshalling doesn't fail
	data := response.Data
	if data == nil {
		data = []byte("{}")
	}

	if err := json.Unmarshal(data, responseBody); err != nil {
		return fmt.Errorf(
			"request %s: unmarshalling `%s` into type %T: %s",
			name,
			data,
			responseBody,
			err,
		)
	}

	return nil
}
