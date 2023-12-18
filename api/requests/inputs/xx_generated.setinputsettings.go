// This file has been automatically generated. Don't edit it.

package inputs

import api "github.com/andreykaipov/goobs/api"

// Represents the request body for the SetInputSettings request.
type SetInputSettingsParams struct {
	// Name of the input to set the settings of
	InputName *string `json:"inputName,omitempty"`

	// Object of settings to apply
	InputSettings map[string]any `json:"inputSettings,omitempty"`

	// True == apply the settings on top of existing ones, False == reset the input to its defaults, then apply
	// settings.
	Overlay *bool `json:"overlay,omitempty"`
}

func NewSetInputSettingsParams() *SetInputSettingsParams {
	return &SetInputSettingsParams{}
}
func (o *SetInputSettingsParams) WithInputName(x string) *SetInputSettingsParams {
	o.InputName = &x
	return o
}
func (o *SetInputSettingsParams) WithInputSettings(x map[string]any) *SetInputSettingsParams {
	o.InputSettings = x
	return o
}
func (o *SetInputSettingsParams) WithOverlay(x bool) *SetInputSettingsParams {
	o.Overlay = &x
	return o
}

// Returns the associated request.
func (o *SetInputSettingsParams) GetRequestName() string {
	return "SetInputSettings"
}

// Represents the response body for the SetInputSettings request.
type SetInputSettingsResponse struct {
	api.ResponseCommon
}

// Sets the settings of an input.
func (c *Client) SetInputSettings(params *SetInputSettingsParams) (*SetInputSettingsResponse, error) {
	data := &SetInputSettingsResponse{}
	return data, c.SendRequest(params, data)
}
