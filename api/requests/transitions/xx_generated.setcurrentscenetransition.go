// This file has been automatically generated. Don't edit it.

package transitions

// Represents the request body for the SetCurrentSceneTransition request.
type SetCurrentSceneTransitionParams struct {
	// Name of the transition to make active
	TransitionName *string `json:"transitionName,omitempty"`
}

func NewSetCurrentSceneTransitionParams() *SetCurrentSceneTransitionParams {
	return &SetCurrentSceneTransitionParams{}
}
func (o *SetCurrentSceneTransitionParams) WithTransitionName(x string) *SetCurrentSceneTransitionParams {
	o.TransitionName = &x
	return o
}

// Returns the associated request.
func (o *SetCurrentSceneTransitionParams) GetRequestName() string {
	return "SetCurrentSceneTransition"
}

// Represents the response body for the SetCurrentSceneTransition request.
type SetCurrentSceneTransitionResponse struct {
	_response
}

/*
Sets the current scene transition.

Small note: While the namespace of scene transitions is generally unique, that uniqueness is not a guarantee as it is with other resources like inputs.
*/
func (c *Client) SetCurrentSceneTransition(
	params *SetCurrentSceneTransitionParams,
) (*SetCurrentSceneTransitionResponse, error) {
	data := &SetCurrentSceneTransitionResponse{}
	return data, c.client.SendRequest(params, data)
}
