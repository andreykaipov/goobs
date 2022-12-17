// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the ToggleOutput request.
type ToggleOutputParams struct {
	// Output name
	OutputName string `json:"outputName,omitempty"`
}

// Returns the associated request.
func (o *ToggleOutputParams) GetRequestName() string {
	return "ToggleOutput"
}

// Represents the response body for the ToggleOutput request.
type ToggleOutputResponse struct {
	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`
}

// Toggles the status of an output.
func (c *Client) ToggleOutput(params *ToggleOutputParams) (*ToggleOutputResponse, error) {
	data := &ToggleOutputResponse{}
	return data, c.SendRequest(params, data)
}
