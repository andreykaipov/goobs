// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the SetOutputSettings request.
type SetOutputSettingsParams struct {
	// Output name
	OutputName string `json:"outputName,omitempty"`

	// Output settings
	OutputSettings interface{} `json:"outputSettings,omitempty"`
}

// Returns the associated request.
func (o *SetOutputSettingsParams) GetRequestName() string {
	return "SetOutputSettings"
}

// Represents the response body for the SetOutputSettings request.
type SetOutputSettingsResponse struct{}

// Sets the settings of an output.
func (c *Client) SetOutputSettings(params *SetOutputSettingsParams) (*SetOutputSettingsResponse, error) {
	data := &SetOutputSettingsResponse{}
	return data, c.SendRequest(params, data)
}
