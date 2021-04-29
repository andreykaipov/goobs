// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetVersionParams represents the params body for the "GetVersion" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVersion.
*/
type GetVersionParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetVersionParams
func (o *GetVersionParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetVersionParams
func (o *GetVersionParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetVersionParams
func (o *GetVersionParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetVersionResponse represents the response body for the "GetVersion" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVersion.
*/
type GetVersionResponse struct {
	requests.Response

	// List of available request types, formatted as a comma-separated list string (e.g. :
	// "Method1,Method2,Method3").
	AvailableRequests string `json:"available-requests"`

	// OBS Studio program version.
	ObsStudioVersion string `json:"obs-studio-version"`

	// obs-websocket plugin version.
	ObsWebsocketVersion string `json:"obs-websocket-version"`

	// OBSRemote compatible API version. Fixed to 1.1 for retrocompatibility.
	Version float64 `json:"version"`
}

// GetMessageID returns the MessageID of GetVersionResponse
func (o *GetVersionResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetVersionResponse
func (o *GetVersionResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetVersionResponse
func (o *GetVersionResponse) GetError() string {
	return o.Error
}

// GetVersion sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetVersion(paramss ...*GetVersionParams) (*GetVersionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetVersionParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetVersion"
	data := &GetVersionResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
