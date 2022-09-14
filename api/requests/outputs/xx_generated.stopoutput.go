// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the StopOutput request.
type StopOutputParams struct {
	// Output name
	OutputName string `json:"outputName,omitempty"`
}

// Returns the associated request.
func (o *StopOutputParams) GetRequestName() string {
	return "StopOutput"
}

// Represents the response body for the StopOutput request.
type StopOutputResponse struct{}

// Stops an output.
func (c *Client) StopOutput(params *StopOutputParams) (*StopOutputResponse, error) {
	data := &StopOutputResponse{}
	return data, c.SendRequest(params, data)
}
