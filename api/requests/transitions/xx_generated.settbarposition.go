// This file has been automatically generated. Don't edit it.

package transitions

// Represents the request body for the SetTBarPosition request.
type SetTBarPositionParams struct {
	// New position
	Position *float64 `json:"position,omitempty"`

	// Whether to release the TBar. Only set `false` if you know that you will be sending another position update
	Release *bool `json:"release,omitempty"`
}

func NewSetTBarPositionParams() *SetTBarPositionParams {
	return &SetTBarPositionParams{}
}
func (o *SetTBarPositionParams) WithPosition(x float64) *SetTBarPositionParams {
	o.Position = &x
	return o
}
func (o *SetTBarPositionParams) WithRelease(x bool) *SetTBarPositionParams {
	o.Release = &x
	return o
}

// Returns the associated request.
func (o *SetTBarPositionParams) GetRequestName() string {
	return "SetTBarPosition"
}

// Represents the response body for the SetTBarPosition request.
type SetTBarPositionResponse struct {
	_response
}

/*
Sets the position of the TBar.

**Very important note**: This will be deprecated and replaced in a future version of obs-websocket.
*/
func (c *Client) SetTBarPosition(params *SetTBarPositionParams) (*SetTBarPositionResponse, error) {
	data := &SetTBarPositionResponse{}
	return data, c.client.SendRequest(params, data)
}
