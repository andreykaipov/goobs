package requests

import (
	"encoding/json"
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

	requestBytes, err := json.Marshal(params)
	if err != nil {
		return err
	}

	if err := conn.WriteMessage(websocket.TextMessage, requestBytes); err != nil {
		return err
	}

	_, responseBytes, err := conn.ReadMessage()
	if err := json.Unmarshal(responseBytes, response); err != nil {
		return err
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
