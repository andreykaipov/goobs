// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetTBarPositionParams represents the params body for the "SetTBarPosition" request.
Set the manual position of the T-Bar (in Studio Mode) to the specified value. Will return an error if OBS is not in studio mode or if the current transition doesn't support T-Bar control.

If your code needs to perform multiple successive T-Bar moves (e.g. : in an animation, or in response to a user moving a T-Bar control in your User Interface), set `release` to false and call `ReleaseTBar` later once the animation/interaction is over.
Since 4.9.0.
*/
type SetTBarPositionParams struct {
	requests.ParamsBasic

	// T-Bar position. This value must be between 0.0 and 1.0.
	Position float64 `json:"position,omitempty"`

	// Whether or not the T-Bar gets released automatically after setting its new position (like a user releasing their
	// mouse button after moving the T-Bar). Call `ReleaseTBar` manually if you set `release` to false. Defaults to
	// true.
	Release bool `json:"release"`
}

// GetSelfName just returns "SetTBarPosition".
func (o *SetTBarPositionParams) GetSelfName() string {
	return "SetTBarPosition"
}

/*
SetTBarPositionResponse represents the response body for the "SetTBarPosition" request.
Set the manual position of the T-Bar (in Studio Mode) to the specified value. Will return an error if OBS is not in studio mode or if the current transition doesn't support T-Bar control.

If your code needs to perform multiple successive T-Bar moves (e.g. : in an animation, or in response to a user moving a T-Bar control in your User Interface), set `release` to false and call `ReleaseTBar` later once the animation/interaction is over.
Since v4.9.0.
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
