// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the StopOutput request.
type StopOutputParams struct {
	// Output name
	OutputName *string `json:"outputName,omitempty"`
}

func NewStopOutputParams() *StopOutputParams {
	return &StopOutputParams{}
}
func (o *StopOutputParams) WithOutputName(x string) *StopOutputParams {
	o.OutputName = &x
	return o
}

// Returns the associated request.
func (o *StopOutputParams) GetRequestName() string {
	return "StopOutput"
}

// Represents the response body for the StopOutput request.
type StopOutputResponse struct {
	_response
}

// Stops an output.
func (c *Client) StopOutput(params *StopOutputParams) (*StopOutputResponse, error) {
	data := &StopOutputResponse{}
	return data, c.client.SendRequest(params, data)
}
