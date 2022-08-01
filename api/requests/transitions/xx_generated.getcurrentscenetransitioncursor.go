// This file has been automatically generated. Don't edit it.

package transitions

/*
GetCurrentSceneTransitionCursorParams represents the params body for the "GetCurrentSceneTransitionCursor" request.
Gets the cursor position of the current scene transition.

Note: `transitionCursor` will return 1.0 when the transition is inactive.
*/
type GetCurrentSceneTransitionCursorParams struct{}

/*
GetCurrentSceneTransitionCursorResponse represents the response body for the "GetCurrentSceneTransitionCursor" request.
Gets the cursor position of the current scene transition.

Note: `transitionCursor` will return 1.0 when the transition is inactive.
*/
type GetCurrentSceneTransitionCursorResponse struct {
	// Cursor position, between 0.0 and 1.0
	TransitionCursor float64 `json:"transitionCursor,omitempty"`
}

// GetCurrentSceneTransitionCursor sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentSceneTransitionCursor(
	paramss ...*GetCurrentSceneTransitionCursorParams,
) (*GetCurrentSceneTransitionCursorResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentSceneTransitionCursorParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("GetCurrentSceneTransitionCursor", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetCurrentSceneTransitionCursorResponse), nil
}
