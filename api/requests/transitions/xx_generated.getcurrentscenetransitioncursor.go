// This file has been automatically generated. Don't edit it.

package transitions

// Represents the request body for the GetCurrentSceneTransitionCursor request.
type GetCurrentSceneTransitionCursorParams struct{}

// Returns the associated request.
func (o *GetCurrentSceneTransitionCursorParams) GetRequestName() string {
	return "GetCurrentSceneTransitionCursor"
}

// Represents the response body for the GetCurrentSceneTransitionCursor request.
type GetCurrentSceneTransitionCursorResponse struct {
	_response

	// Cursor position, between 0.0 and 1.0
	TransitionCursor float64 `json:"transitionCursor,omitempty"`
}

/*
Gets the cursor position of the current scene transition.

Note: `transitionCursor` will return 1.0 when the transition is inactive.
*/
func (c *Client) GetCurrentSceneTransitionCursor(
	paramss ...*GetCurrentSceneTransitionCursorParams,
) (*GetCurrentSceneTransitionCursorResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentSceneTransitionCursorParams{{}}
	}
	params := paramss[0]
	data := &GetCurrentSceneTransitionCursorResponse{}
	return data, c.client.SendRequest(params, data)
}
