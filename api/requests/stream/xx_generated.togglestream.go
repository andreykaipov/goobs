// This file has been automatically generated. Don't edit it.

package stream

// Represents the request body for the ToggleStream request.
type ToggleStreamParams struct{}

// Returns the associated request.
func (o *ToggleStreamParams) GetRequestName() string {
	return "ToggleStream"
}

// Represents the response body for the ToggleStream request.
type ToggleStreamResponse struct {
	_response

	// New state of the stream output
	OutputActive bool `json:"outputActive,omitempty"`
}

// Toggles the status of the stream output.
func (c *Client) ToggleStream(paramss ...*ToggleStreamParams) (*ToggleStreamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleStreamParams{{}}
	}
	params := paramss[0]
	data := &ToggleStreamResponse{}
	return data, c.client.SendRequest(params, data)
}
