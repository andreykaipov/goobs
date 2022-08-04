// This file has been automatically generated. Don't edit it.

package ui

/*
SetStudioModeEnabledParams represents the params body for the "SetStudioModeEnabled" request.
Enables or disables studio mode
*/
type SetStudioModeEnabledParams struct {
	// True == Enabled, False == Disabled
	StudioModeEnabled *bool `json:"studioModeEnabled,omitempty"`
}

/*
SetStudioModeEnabledResponse represents the response body for the "SetStudioModeEnabled" request.
Enables or disables studio mode
*/
type SetStudioModeEnabledResponse struct{}

// SetStudioModeEnabled sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetStudioModeEnabled(params *SetStudioModeEnabledParams) (*SetStudioModeEnabledResponse, error) {
	resp, err := c.SendRequest("SetStudioModeEnabled", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetStudioModeEnabledResponse), nil
}
