// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetTBarPositionParams represents the params body for the "SetTBarPosition" request.
Sets the position of the TBar.

**Very important note**: This will be deprecated and replaced in a future version of obs-websocket.
*/
type SetTBarPositionParams struct {
	requests.ParamsBasic

	// New position
	Position float64 `json:"position,omitempty"`

	// Whether to release the TBar. Only set `false` if you know that you will be sending another position update
	Release bool `json:"release,omitempty"`
}

// GetSelfName just returns "SetTBarPosition".
func (o *SetTBarPositionParams) GetSelfName() string {
	return "SetTBarPosition"
}

/*
SetTBarPositionResponse represents the response body for the "SetTBarPosition" request.
Sets the position of the TBar.

**Very important note**: This will be deprecated and replaced in a future version of obs-websocket.
*/
type SetTBarPositionResponse struct {
	requests.ResponseBasic
}

// SetTBarPosition sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetTBarPosition(params *SetTBarPositionParams) (*SetTBarPositionResponse, error) {
	data := &SetTBarPositionResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
