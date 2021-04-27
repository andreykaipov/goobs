// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetCurrentTransitionParams represents the params body for the "GetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentTransition.
*/
type GetCurrentTransitionParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetCurrentTransitionParams
func (o *GetCurrentTransitionParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetCurrentTransitionParams
func (o *GetCurrentTransitionParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetCurrentTransitionParams
func (o *GetCurrentTransitionParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetCurrentTransitionResponse represents the response body for the "GetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentTransition.
*/
type GetCurrentTransitionResponse struct {
	requests.Response

	// Transition duration (in milliseconds) if supported by the transition.
	Duration int `json:"duration"`

	// Name of the selected transition.
	Name string `json:"name"`
}

// GetMessageID returns the MessageID of GetCurrentTransitionResponse
func (o *GetCurrentTransitionResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetCurrentTransitionResponse
func (o *GetCurrentTransitionResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetCurrentTransitionResponse
func (o *GetCurrentTransitionResponse) GetError() string {
	return o.Error
}

// GetCurrentTransition sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentTransition(
	paramss ...*GetCurrentTransitionParams,
) (*GetCurrentTransitionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentTransitionParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetCurrentTransition"
	data := &GetCurrentTransitionResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
