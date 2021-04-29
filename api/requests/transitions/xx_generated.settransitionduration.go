// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetTransitionDurationParams represents the params body for the "SetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTransitionDuration.
*/
type SetTransitionDurationParams struct {
	requests.Params

	// Desired duration of the transition (in milliseconds).
	Duration int `json:"duration"`
}

// GetRequestType returns the RequestType of SetTransitionDurationParams
func (o *SetTransitionDurationParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetTransitionDurationParams
func (o *SetTransitionDurationParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetTransitionDurationParams
func (o *SetTransitionDurationParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetTransitionDurationResponse represents the response body for the "SetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTransitionDuration.
*/
type SetTransitionDurationResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetTransitionDurationResponse
func (o *SetTransitionDurationResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetTransitionDurationResponse
func (o *SetTransitionDurationResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetTransitionDurationResponse
func (o *SetTransitionDurationResponse) GetError() string {
	return o.Error
}

// SetTransitionDuration sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetTransitionDuration(
	params *SetTransitionDurationParams,
) (*SetTransitionDurationResponse, error) {
	params.RequestType = "SetTransitionDuration"
	data := &SetTransitionDurationResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
