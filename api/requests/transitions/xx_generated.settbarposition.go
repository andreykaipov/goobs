// This file has been automatically generated. Don't edit it.

package transitions

// Represents the request body for the SetTBarPosition request.
type SetTBarPositionParams struct {
	// New position
	Position float64 `json:"position,omitempty"`

	// Whether to release the TBar. Only set `false` if you know that you will be sending another position update
	Release *bool `json:"release,omitempty"`
}

// Returns the associated request.
func (o *SetTBarPositionParams) GetRequestName() string {
	return "SetTBarPosition"
}

// Represents the response body for the SetTBarPosition request.
type SetTBarPositionResponse struct{}

/*
Sets the position of the TBar.

**Very important note**: This will be deprecated and replaced in a future version of obs-websocket.
*/
func (c *Client) SetTBarPosition(params *SetTBarPositionParams) (*SetTBarPositionResponse, error) {
	data := &SetTBarPositionResponse{}
	return data, c.SendRequest(params, data)
}
