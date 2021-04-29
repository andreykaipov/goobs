// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
EnableStudioModeParams represents the params body for the "EnableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#EnableStudioMode.
*/
type EnableStudioModeParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of EnableStudioModeParams
func (o *EnableStudioModeParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of EnableStudioModeParams
func (o *EnableStudioModeParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on EnableStudioModeParams
func (o *EnableStudioModeParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
EnableStudioModeResponse represents the response body for the "EnableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#EnableStudioMode.
*/
type EnableStudioModeResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of EnableStudioModeResponse
func (o *EnableStudioModeResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of EnableStudioModeResponse
func (o *EnableStudioModeResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of EnableStudioModeResponse
func (o *EnableStudioModeResponse) GetError() string {
	return o.Error
}

// EnableStudioMode sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) EnableStudioMode(
	paramss ...*EnableStudioModeParams,
) (*EnableStudioModeResponse, error) {
	if len(paramss) == 0 {
		paramss = []*EnableStudioModeParams{{}}
	}
	params := paramss[0]
	params.RequestType = "EnableStudioMode"
	data := &EnableStudioModeResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
