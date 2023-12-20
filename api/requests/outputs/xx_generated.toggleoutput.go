// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the ToggleOutput request.
type ToggleOutputParams struct {
	// Output name
	OutputName *string `json:"outputName,omitempty"`
}

func NewToggleOutputParams() *ToggleOutputParams {
	return &ToggleOutputParams{}
}
func (o *ToggleOutputParams) WithOutputName(x string) *ToggleOutputParams {
	o.OutputName = &x
	return o
}

// Returns the associated request.
func (o *ToggleOutputParams) GetRequestName() string {
	return "ToggleOutput"
}

// Represents the response body for the ToggleOutput request.
type ToggleOutputResponse struct {
	_response

	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`
}

// Toggles the status of an output.
func (c *Client) ToggleOutput(params *ToggleOutputParams) (*ToggleOutputResponse, error) {
	data := &ToggleOutputResponse{}
	return data, c.client.SendRequest(params, data)
}
