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

type Params struct {
	paramsBehavior
	RequestType string `json:"request-type"`
	MessageID   string `json:"message-id"`
}

func (o *Params) GetRequestType() string {
	return o.RequestType
}
func (o *Params) GetMessageID() string {
	return o.MessageID
}
func (o *Params) SetMessageID(x string) {
	o.MessageID = x
}

type responseBehavior interface {
	GetMessageID() string
	GetStatus() string
	GetError() string
}

type Response struct {
	responseBehavior
	MessageID string `json:"message-id"`
	Status    string `json:"status"`
	Error     string `json:"error"`
}

func (o *Response) GetMessageID() string {
	return o.MessageID
}
func (o *Response) GetStatus() string {
	return o.Status
}
func (o *Response) GetError() string {
	return o.Error
}

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
