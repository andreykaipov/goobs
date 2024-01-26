// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the SetInputSettings request.
type SetInputSettingsParams struct {
	// Name of the input to set the settings of
	InputName *string `json:"inputName,omitempty"`

	// Object of settings to apply
	InputSettings map[string]any `json:"inputSettings,omitempty"`

	// UUID of the input to set the settings of
	InputUuid *string `json:"inputUuid,omitempty"`

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
func (o *SetInputSettingsParams) WithInputUuid(x string) *SetInputSettingsParams {
	o.InputUuid = &x
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
	_response
}

// Sets the settings of an input.
func (c *Client) SetInputSettings(params *SetInputSettingsParams) (*SetInputSettingsResponse, error) {
	data := &SetInputSettingsResponse{}
	return data, c.client.SendRequest(params, data)
}
