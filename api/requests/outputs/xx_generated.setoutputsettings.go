// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the SetOutputSettings request.
type SetOutputSettingsParams struct {
	// Output name
	OutputName *string `json:"outputName,omitempty"`

	// Output settings
	OutputSettings map[string]any `json:"outputSettings,omitempty"`
}

func NewSetOutputSettingsParams() *SetOutputSettingsParams {
	return &SetOutputSettingsParams{}
}
func (o *SetOutputSettingsParams) WithOutputName(x string) *SetOutputSettingsParams {
	o.OutputName = &x
	return o
}
func (o *SetOutputSettingsParams) WithOutputSettings(x map[string]any) *SetOutputSettingsParams {
	o.OutputSettings = x
	return o
}

// Returns the associated request.
func (o *SetOutputSettingsParams) GetRequestName() string {
	return "SetOutputSettings"
}

// Represents the response body for the SetOutputSettings request.
type SetOutputSettingsResponse struct {
	_response
}

// Sets the settings of an output.
func (c *Client) SetOutputSettings(params *SetOutputSettingsParams) (*SetOutputSettingsResponse, error) {
	data := &SetOutputSettingsResponse{}
	return data, c.client.SendRequest(params, data)
}
