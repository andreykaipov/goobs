// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetTransitionDurationParams represents the params body for the "GetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionDuration.
*/
type GetTransitionDurationParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetTransitionDurationParams
func (o *GetTransitionDurationParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetTransitionDurationParams
func (o *GetTransitionDurationParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetTransitionDurationParams
func (o *GetTransitionDurationParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetTransitionDurationResponse represents the response body for the "GetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionDuration.
*/
type GetTransitionDurationResponse struct {
	requests.Response

	// Duration of the current transition (in milliseconds).
	TransitionDuration int `json:"transition-duration"`
}

// GetMessageID returns the MessageID of GetTransitionDurationResponse
func (o *GetTransitionDurationResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetTransitionDurationResponse
func (o *GetTransitionDurationResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetTransitionDurationResponse
func (o *GetTransitionDurationResponse) GetError() string {
	return o.Error
}

// GetTransitionDuration sends the corresponding request to the connected OBS WebSockets server.
// Note the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetTransitionDuration(
	paramss ...*GetTransitionDurationParams,
) (*GetTransitionDurationResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetTransitionDurationParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetTransitionDuration"
	data := &GetTransitionDurationResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
