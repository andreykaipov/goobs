// This file has been automatically generated. Don't edit it.

package ui

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetStudioModeEnabledParams represents the params body for the "GetStudioModeEnabled" request.
Gets whether studio is enabled.
*/
type GetStudioModeEnabledParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetStudioModeEnabled".
func (o *GetStudioModeEnabledParams) GetSelfName() string {
	return "GetStudioModeEnabled"
}

/*
GetStudioModeEnabledResponse represents the response body for the "GetStudioModeEnabled" request.
Gets whether studio is enabled.
*/
type GetStudioModeEnabledResponse struct {
	requests.ResponseBasic

	// Whether studio mode is enabled
	StudioModeEnabled bool `json:"studioModeEnabled,omitempty"`
}

// GetStudioModeEnabled sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetStudioModeEnabled(paramss ...*GetStudioModeEnabledParams) (*GetStudioModeEnabledResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStudioModeEnabledParams{{}}
	}
	params := paramss[0]
	data := &GetStudioModeEnabledResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
