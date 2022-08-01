// This file has been automatically generated. Don't edit it.

package transitions

/*
SetTBarPositionParams represents the params body for the "SetTBarPosition" request.
Sets the position of the TBar.

**Very important note**: This will be deprecated and replaced in a future version of obs-websocket.
*/
type SetTBarPositionParams struct {
	// New position
	Position float64 `json:"position,omitempty"`

	// Whether to release the TBar. Only set `false` if you know that you will be sending another position update
	Release bool `json:"release,omitempty"`
}

/*
SetTBarPositionResponse represents the response body for the "SetTBarPosition" request.
Sets the position of the TBar.

**Very important note**: This will be deprecated and replaced in a future version of obs-websocket.
*/
type SetTBarPositionResponse struct{}

// SetTBarPosition sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetTBarPosition(params *SetTBarPositionParams) (*SetTBarPositionResponse, error) {
	resp, err := c.SendRequest("SetTBarPosition", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetTBarPositionResponse), nil
}
