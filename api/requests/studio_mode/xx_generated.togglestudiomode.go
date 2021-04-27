// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ToggleStudioModeParams represents the params body for the "ToggleStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleStudioMode.
*/
type ToggleStudioModeParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of ToggleStudioModeParams
func (o *ToggleStudioModeParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of ToggleStudioModeParams
func (o *ToggleStudioModeParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on ToggleStudioModeParams
func (o *ToggleStudioModeParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ToggleStudioModeResponse represents the response body for the "ToggleStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleStudioMode.
*/
type ToggleStudioModeResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of ToggleStudioModeResponse
func (o *ToggleStudioModeResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of ToggleStudioModeResponse
func (o *ToggleStudioModeResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of ToggleStudioModeResponse
func (o *ToggleStudioModeResponse) GetError() string {
	return o.Error
}

// ToggleStudioMode sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) ToggleStudioMode(
	paramss ...*ToggleStudioModeParams,
) (*ToggleStudioModeResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleStudioModeParams{{}}
	}
	params := paramss[0]
	params.RequestType = "ToggleStudioMode"
	data := &ToggleStudioModeResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
