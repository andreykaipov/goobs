// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the GetOutputSettings request.
type GetOutputSettingsParams struct {
	// Output name
	OutputName *string `json:"outputName,omitempty"`
}

func NewGetOutputSettingsParams() *GetOutputSettingsParams {
	return &GetOutputSettingsParams{}
}
func (o *GetOutputSettingsParams) WithOutputName(x string) *GetOutputSettingsParams {
	o.OutputName = &x
	return o
}

// Returns the associated request.
func (o *GetOutputSettingsParams) GetRequestName() string {
	return "GetOutputSettings"
}

// Represents the response body for the GetOutputSettings request.
type GetOutputSettingsResponse struct {
	_response

	// Output settings
	OutputSettings map[string]any `json:"outputSettings,omitempty"`
}

// Gets the settings of an output.
func (c *Client) GetOutputSettings(params *GetOutputSettingsParams) (*GetOutputSettingsResponse, error) {
	data := &GetOutputSettingsResponse{}
	return data, c.client.SendRequest(params, data)
}
