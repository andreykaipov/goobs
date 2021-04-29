// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
TransitionToProgramParams represents the params body for the "TransitionToProgram" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#TransitionToProgram.
*/
type TransitionToProgramParams struct {
	requests.Params

	WithTransition struct {
		// Transition duration (in milliseconds).
		Duration int `json:"duration"`

		// Name of the transition.
		Name string `json:"name"`
	} `json:"with-transition"`
}

// GetRequestType returns the RequestType of TransitionToProgramParams
func (o *TransitionToProgramParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of TransitionToProgramParams
func (o *TransitionToProgramParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on TransitionToProgramParams
func (o *TransitionToProgramParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
TransitionToProgramResponse represents the response body for the "TransitionToProgram" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#TransitionToProgram.
*/
type TransitionToProgramResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of TransitionToProgramResponse
func (o *TransitionToProgramResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of TransitionToProgramResponse
func (o *TransitionToProgramResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of TransitionToProgramResponse
func (o *TransitionToProgramResponse) GetError() string {
	return o.Error
}

// TransitionToProgram sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) TransitionToProgram(
	params *TransitionToProgramParams,
) (*TransitionToProgramResponse, error) {
	params.RequestType = "TransitionToProgram"
	data := &TransitionToProgramResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
