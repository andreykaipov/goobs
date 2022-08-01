// This file has been automatically generated. Don't edit it.

package stream

/*
ToggleStreamParams represents the params body for the "ToggleStream" request.
Toggles the status of the stream output.
*/
type ToggleStreamParams struct{}

/*
ToggleStreamResponse represents the response body for the "ToggleStream" request.
Toggles the status of the stream output.
*/
type ToggleStreamResponse struct {
	// New state of the stream output
	OutputActive bool `json:"outputActive,omitempty"`
}

// ToggleStream sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) ToggleStream(paramss ...*ToggleStreamParams) (*ToggleStreamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleStreamParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("ToggleStream", params)
	if err != nil {
		return nil, err
	}
	return resp.(*ToggleStreamResponse), nil
}
