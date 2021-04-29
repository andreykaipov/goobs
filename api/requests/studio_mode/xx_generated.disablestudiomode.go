// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
DisableStudioModeParams represents the params body for the "DisableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DisableStudioMode.
*/
type DisableStudioModeParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of DisableStudioModeParams
func (o *DisableStudioModeParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of DisableStudioModeParams
func (o *DisableStudioModeParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on DisableStudioModeParams
func (o *DisableStudioModeParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
DisableStudioModeResponse represents the response body for the "DisableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DisableStudioMode.
*/
type DisableStudioModeResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of DisableStudioModeResponse
func (o *DisableStudioModeResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of DisableStudioModeResponse
func (o *DisableStudioModeResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of DisableStudioModeResponse
func (o *DisableStudioModeResponse) GetError() string {
	return o.Error
}

// DisableStudioMode sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) DisableStudioMode(
	paramss ...*DisableStudioModeParams,
) (*DisableStudioModeResponse, error) {
	if len(paramss) == 0 {
		paramss = []*DisableStudioModeParams{{}}
	}
	params := paramss[0]
	params.RequestType = "DisableStudioMode"
	data := &DisableStudioModeResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
