// This file has been automatically generated. Don't edit it.

package inputs

/*
SetInputSettingsParams represents the params body for the "SetInputSettings" request.
Sets the settings of an input.
*/
type SetInputSettingsParams struct {
	// Name of the input to set the settings of
	InputName string `json:"inputName,omitempty"`

	// Object of settings to apply
	InputSettings interface{} `json:"inputSettings,omitempty"`

	// True == apply the settings on top of existing ones, False == reset the input to its defaults, then apply
	// settings.
	Overlay *bool `json:"overlay,omitempty"`
}

/*
SetInputSettingsResponse represents the response body for the "SetInputSettings" request.
Sets the settings of an input.
*/
type SetInputSettingsResponse struct{}

// SetInputSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputSettings(params *SetInputSettingsParams) (*SetInputSettingsResponse, error) {
	resp, err := c.SendRequest("SetInputSettings", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetInputSettingsResponse), nil
}
