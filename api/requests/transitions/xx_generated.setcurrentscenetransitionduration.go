// This file has been automatically generated. Don't edit it.

package transitions

/*
SetCurrentSceneTransitionDurationParams represents the params body for the "SetCurrentSceneTransitionDuration" request.
Sets the duration of the current scene transition, if it is not fixed.
*/
type SetCurrentSceneTransitionDurationParams struct {
	// Duration in milliseconds
	TransitionDuration float64 `json:"transitionDuration,omitempty"`
}

/*
SetCurrentSceneTransitionDurationResponse represents the response body for the "SetCurrentSceneTransitionDuration" request.
Sets the duration of the current scene transition, if it is not fixed.
*/
type SetCurrentSceneTransitionDurationResponse struct{}

// SetCurrentSceneTransitionDuration sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentSceneTransitionDuration(
	params *SetCurrentSceneTransitionDurationParams,
) (*SetCurrentSceneTransitionDurationResponse, error) {
	resp, err := c.SendRequest("SetCurrentSceneTransitionDuration", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetCurrentSceneTransitionDurationResponse), nil
}
