// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetInputSettingsParams represents the params body for the "SetInputSettings" request.
Sets the settings of an input.
*/
type SetInputSettingsParams struct {
	requests.ParamsBasic

	// Name of the input to set the settings of
	InputName string `json:"inputName,omitempty"`

	// Object of settings to apply
	InputSettings interface{} `json:"inputSettings,omitempty"`

	// True == apply the settings on top of existing ones, False == reset the input to its defaults, then apply
	// settings.
	Overlay bool `json:"overlay,omitempty"`
}

// GetSelfName just returns "SetInputSettings".
func (o *SetInputSettingsParams) GetSelfName() string {
	return "SetInputSettings"
}

/*
SetInputSettingsResponse represents the response body for the "SetInputSettings" request.
Sets the settings of an input.
*/
type SetInputSettingsResponse struct {
	requests.ResponseBasic
}

// SetInputSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputSettings(params *SetInputSettingsParams) (*SetInputSettingsResponse, error) {
	data := &SetInputSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
