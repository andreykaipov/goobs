// This file has been automatically generated. Don't edit it.

package transitions

// Represents the request body for the SetCurrentSceneTransitionDuration request.
type SetCurrentSceneTransitionDurationParams struct {
	// Duration in milliseconds
	TransitionDuration *float64 `json:"transitionDuration,omitempty"`
}

func NewSetCurrentSceneTransitionDurationParams() *SetCurrentSceneTransitionDurationParams {
	return &SetCurrentSceneTransitionDurationParams{}
}

func (o *SetCurrentSceneTransitionDurationParams) WithTransitionDuration(
	x float64,
) *SetCurrentSceneTransitionDurationParams {
	o.TransitionDuration = &x
	return o
}

// Returns the associated request.
func (o *SetCurrentSceneTransitionDurationParams) GetRequestName() string {
	return "SetCurrentSceneTransitionDuration"
}

// Represents the response body for the SetCurrentSceneTransitionDuration request.
type SetCurrentSceneTransitionDurationResponse struct {
	_response
}

// Sets the duration of the current scene transition, if it is not fixed.
func (c *Client) SetCurrentSceneTransitionDuration(
	params *SetCurrentSceneTransitionDurationParams,
) (*SetCurrentSceneTransitionDurationResponse, error) {
	data := &SetCurrentSceneTransitionDurationResponse{}
	return data, c.client.SendRequest(params, data)
}
