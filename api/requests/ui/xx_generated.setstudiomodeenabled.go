// This file has been automatically generated. Don't edit it.

package ui

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetStudioModeEnabledParams represents the params body for the "SetStudioModeEnabled" request.
Enables or disables studio mode
*/
type SetStudioModeEnabledParams struct {
	requests.ParamsBasic

	// True == Enabled, False == Disabled
	StudioModeEnabled bool `json:"studioModeEnabled,omitempty"`
}

// GetSelfName just returns "SetStudioModeEnabled".
func (o *SetStudioModeEnabledParams) GetSelfName() string {
	return "SetStudioModeEnabled"
}

/*
SetStudioModeEnabledResponse represents the response body for the "SetStudioModeEnabled" request.
Enables or disables studio mode
*/
type SetStudioModeEnabledResponse struct {
	requests.ResponseBasic
}

// SetStudioModeEnabled sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetStudioModeEnabled(params *SetStudioModeEnabledParams) (*SetStudioModeEnabledResponse, error) {
	data := &SetStudioModeEnabledResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
