// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetVersionParams represents the params body for the "GetVersion" request.
Returns the latest version of the plugin and the API.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#GetVersion.
*/
type GetVersionParams struct {
	requests.ParamsBasic
}

// Name just returns "GetVersion".
func (o *GetVersionParams) Name() string {
	return "GetVersion"
}

/*
GetVersionResponse represents the response body for the "GetVersion" request.
Returns the latest version of the plugin and the API.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#GetVersion.
*/
type GetVersionResponse struct {
	requests.ResponseBasic

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

// GetVersion sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetVersion(paramss ...*GetVersionParams) (*GetVersionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetVersionParams{{}}
	}
	params := paramss[0]
	data := &GetVersionResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
