// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetStudioModeStatusParams represents the params body for the "GetStudioModeStatus" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStudioModeStatus.
*/
type GetStudioModeStatusParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetStudioModeStatusParams
func (o *GetStudioModeStatusParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetStudioModeStatusParams
func (o *GetStudioModeStatusParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetStudioModeStatusParams
func (o *GetStudioModeStatusParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetStudioModeStatusResponse represents the response body for the "GetStudioModeStatus" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStudioModeStatus.
*/
type GetStudioModeStatusResponse struct {
	requests.Response

	// Indicates if Studio Mode is enabled.
	StudioMode bool `json:"studio-mode"`
}

// GetMessageID returns the MessageID of GetStudioModeStatusResponse
func (o *GetStudioModeStatusResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetStudioModeStatusResponse
func (o *GetStudioModeStatusResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetStudioModeStatusResponse
func (o *GetStudioModeStatusResponse) GetError() string {
	return o.Error
}

// GetStudioModeStatus sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetStudioModeStatus(
	paramss ...*GetStudioModeStatusParams,
) (*GetStudioModeStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStudioModeStatusParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetStudioModeStatus"
	data := &GetStudioModeStatusResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
