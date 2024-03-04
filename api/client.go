// Package api is the intermediary API between the top-level goobs client and
// the category-level subclients.
//
// Nothing in this package should be of interest to a user.
package api

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/andreykaipov/goobs/api/opcodes"
	"github.com/andreykaipov/goobs/api/requests"
	uuid "github.com/nu7hatch/gouuid"
)

type Params interface{ GetRequestName() string }

type ResponseCommon struct{ raw json.RawMessage }

func (o *ResponseCommon) setRaw(raw json.RawMessage) { o.raw = raw }
func (o *ResponseCommon) GetRaw() json.RawMessage    { return o.raw }

type Response interface {
	setRaw(json.RawMessage)
	GetRaw() json.RawMessage
}

// Client represents a minimal client to the OBS websocket server.
type Client struct {
	// The time we're willing to wait to receive a response from the server.
	ResponseTimeout time.Duration

	// This client sends raw opcodes it got from the server to this channel.
	Opcodes chan opcodes.Opcode

	// Once the top-level has parsed the raw opcode, it sends the response
	// to this channel.
	IncomingResponses chan *opcodes.RequestResponse

	// Ya like logs?
	Log Logger

	Disconnected chan bool

	mutex sync.Mutex
}

// SendRequest abstracts the logic every subclient uses to send a request and
// receive the corresponding response.
//
// To get the response for a sent request, we simply read the next response
// off our incoming responses channel. This works fine in a single-threaded
// context, and the message IDs of both the sent request and response should
// match.
//
// A request ID and response ID mismatch could happen if the server processes
// requests in a different order it received them (e.g. we should 1, then 2; but
// it processes 2, and then 1). In this case there'll be an error, so note the
// mutex lock and deferred unlock to prevent this from happening.
//
// It should be noted multiple connections to the server are totally fine.
// Phrased differently, mesasge IDs are unique per client. Moreover, events will
// be broadcast to every client.
func (c *Client) SendRequest(requestBody Params, responseBody Response) error {
	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	name := requestBody.GetRequestName()
	id := uid.String()

	c.Log.Printf("[TRACE] Sending %s Request with ID %s", name, id)

	c.mutex.Lock()
	defer c.mutex.Unlock()

	select {
	case <-c.Disconnected:
		return fmt.Errorf("request %s: client already disconnected", name)
	default:
		c.Opcodes <- &opcodes.Request{
			Type: name,
			ID:   id,
			Data: requestBody,
		}
	}

	var response *opcodes.RequestResponse

	timer := time.NewTimer(c.ResponseTimeout * time.Millisecond)
	defer timer.Stop()
	select {
	case response = <-c.IncomingResponses:
	case <-timer.C:
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

	responseBody.setRaw(data)

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
