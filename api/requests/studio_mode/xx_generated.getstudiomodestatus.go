// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetStudioModeStatusParams represents the params body for the "GetStudioModeStatus" request.
Indicates if Studio Mode is currently enabled.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#GetStudioModeStatus.
*/
type GetStudioModeStatusParams struct {
	requests.ParamsBasic
}

// Name just returns "GetStudioModeStatus".
func (o *GetStudioModeStatusParams) Name() string {
	return "GetStudioModeStatus"
}

/*
GetStudioModeStatusResponse represents the response body for the "GetStudioModeStatus" request.
Indicates if Studio Mode is currently enabled.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#GetStudioModeStatus.
*/
type GetStudioModeStatusResponse struct {
	requests.ResponseBasic

	// Indicates if Studio Mode is enabled.
	StudioMode bool `json:"studio-mode"`
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
	data := &GetStudioModeStatusResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
