// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the GetOutputSettings request.
type GetOutputSettingsParams struct {
	// Output name
	OutputName string `json:"outputName,omitempty"`
}

// Returns the associated request.
func (o *GetOutputSettingsParams) GetRequestName() string {
	return "GetOutputSettings"
}

// Represents the response body for the GetOutputSettings request.
type GetOutputSettingsResponse struct {
	// Output settings
	OutputSettings interface{} `json:"outputSettings,omitempty"`
}

// Gets the settings of an output.
func (c *Client) GetOutputSettings(params *GetOutputSettingsParams) (*GetOutputSettingsResponse, error) {
	data := &GetOutputSettingsResponse{}
	return data, c.SendRequest(params, data)
}
