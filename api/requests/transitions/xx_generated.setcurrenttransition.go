// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentTransitionParams represents the params body for the "SetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentTransition.
*/
type SetCurrentTransitionParams struct {
	requests.Params

	// The name of the transition.
	TransitionName string `json:"transition-name"`
}

// GetRequestType returns the RequestType of SetCurrentTransitionParams
func (o *SetCurrentTransitionParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetCurrentTransitionParams
func (o *SetCurrentTransitionParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetCurrentTransitionParams
func (o *SetCurrentTransitionParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetCurrentTransitionResponse represents the response body for the "SetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentTransition.
*/
type SetCurrentTransitionResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetCurrentTransitionResponse
func (o *SetCurrentTransitionResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetCurrentTransitionResponse
func (o *SetCurrentTransitionResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetCurrentTransitionResponse
func (o *SetCurrentTransitionResponse) GetError() string {
	return o.Error
}

// SetCurrentTransition sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentTransition(
	params *SetCurrentTransitionParams,
) (*SetCurrentTransitionResponse, error) {
	params.RequestType = "SetCurrentTransition"
	data := &SetCurrentTransitionResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
