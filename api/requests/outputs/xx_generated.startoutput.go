// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the StartOutput request.
type StartOutputParams struct {
	// Output name
	OutputName string `json:"outputName,omitempty"`
}

// Returns the associated request.
func (o *StartOutputParams) GetRequestName() string {
	return "StartOutput"
}

// Represents the response body for the StartOutput request.
type StartOutputResponse struct{}

// Starts an output.
func (c *Client) StartOutput(params *StartOutputParams) (*StartOutputResponse, error) {
	data := &StartOutputResponse{}
	return data, c.SendRequest(params, data)
}
