// This file has been automatically generated. Don't edit it.

package ui

import api "github.com/andreykaipov/goobs/api"

// Represents the request body for the SetStudioModeEnabled request.
type SetStudioModeEnabledParams struct {
	// True == Enabled, False == Disabled
	StudioModeEnabled *bool `json:"studioModeEnabled,omitempty"`
}

func NewSetStudioModeEnabledParams() *SetStudioModeEnabledParams {
	return &SetStudioModeEnabledParams{}
}
func (o *SetStudioModeEnabledParams) WithStudioModeEnabled(x bool) *SetStudioModeEnabledParams {
	o.StudioModeEnabled = &x
	return o
}

// Returns the associated request.
func (o *SetStudioModeEnabledParams) GetRequestName() string {
	return "SetStudioModeEnabled"
}

// Represents the response body for the SetStudioModeEnabled request.
type SetStudioModeEnabledResponse struct {
	api.ResponseCommon
}

// Enables or disables studio mode
func (c *Client) SetStudioModeEnabled(params *SetStudioModeEnabledParams) (*SetStudioModeEnabledResponse, error) {
	data := &SetStudioModeEnabledResponse{}
	return data, c.SendRequest(params, data)
}
