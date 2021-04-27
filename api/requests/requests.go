package requests

import (
	"fmt"

	"github.com/gorilla/websocket"
	uuid "github.com/nu7hatch/gouuid"
)

type paramsBehavior interface {
	GetRequestType() string
	GetMessageID() string
	SetMessageID(string)
}

// Params represents common parameters for any request.
type Params struct {
	RequestType string `json:"request-type"`
	MessageID   string `json:"message-id"`
}

type responseBehavior interface {
	GetMessageID() string
	GetStatus() string
	GetError() string
}

// Response represents common fields on any returned response.
type Response struct {
	MessageID string `json:"message-id"`
	Status    string `json:"status"`
	Error     string `json:"error"`
}

// WriteMessage abstracts the logic every subclient uses to send a request and
// receive the corresponding response.
func WriteMessage(conn *websocket.Conn, params paramsBehavior, response responseBehavior) error {
	if params.GetMessageID() == "" {
		id, err := uuid.NewV4()
		if err != nil {
			return err
		}

		params.SetMessageID(id.String())
	}

	if err := conn.WriteJSON(params); err != nil {
		return err
	}

	// Some requests respond with some empty response, so keep reading until
	// our message IDs match. Clearly, this isn't safe against concurrency,
	// but gorilla/websocket doesn't handle concurrency either, so this is
	// good enough.
	//
	// TODO: Interestingly, it does seem thread-safe if I use a different
	// connection, so message IDs must be unique per client, in that
	// connection A won't get a response from OBS for a request from
	// connection B. Typing that out, it seems pretty apparent, but I'd like
	// to do some more testing.
	for {
		if err := conn.ReadJSON(response); err != nil {
			return err
		}

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
