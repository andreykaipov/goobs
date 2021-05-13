package requests

import (
	"fmt"

	"github.com/gorilla/websocket"
	uuid "github.com/nu7hatch/gouuid"
)

// Params describes the behavior for any params-like object. Used to abstract
// the functionality of any request that embeds ParamsBasic within their fields.
type Params interface {
	GetRequestType() string
	SetRequestType(string)
	GetMessageID() string
	SetMessageID(string)

	// The name of the actual request, i.e. "Abc" for "AbcParams". Used to
	// set the RequestType.
	Name() string
}

// ParamsBasic represents common parameters for any request.
type ParamsBasic struct {
	RequestType string `json:"request-type"`
	MessageID   string `json:"message-id"`
}

// GetRequestType does what it says.
func (o *ParamsBasic) GetRequestType() string {
	return o.RequestType
}

// SetRequestType does what it says.
func (o *ParamsBasic) SetRequestType(x string) {
	o.RequestType = x
}

// GetMessageID does what it says.
func (o *ParamsBasic) GetMessageID() string {
	return o.MessageID
}

// SetMessageID does what it says.
func (o *ParamsBasic) SetMessageID(x string) {
	o.MessageID = x
}

// Response describes the behavior for any response-like object. Used to
// abstract the functionality of any request's response that embeds
// ResponseBasic within their fields.
type Response interface {
	GetMessageID() string
	GetStatus() string
	GetError() string
}

// ResponseBasic represents common fields on any returned response.
type ResponseBasic struct {
	MessageID string `json:"message-id"`
	Status    string `json:"status"`
	Error     string `json:"error"`
}

// GetMessageID does what it says.
func (o *ResponseBasic) GetMessageID() string {
	return o.MessageID
}

// GetStatus does what it says.
func (o *ResponseBasic) GetStatus() string {
	return o.Status
}

// GetError does what it says.
func (o *ResponseBasic) GetError() string {
	return o.Error
}

// WriteMessage abstracts the logic every subclient uses to send a request and
// receive the corresponding response.
func WriteMessage(conn *websocket.Conn, params Params, response Response) error {
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

	if err := conn.WriteJSON(params); err != nil {
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
