// This file has been automatically generated. Don't edit it.

package transitions

/*
SetCurrentSceneTransitionParams represents the params body for the "SetCurrentSceneTransition" request.
Sets the current scene transition.

Small note: While the namespace of scene transitions is generally unique, that uniqueness is not a guarantee as it is with other resources like inputs.
*/
type SetCurrentSceneTransitionParams struct {
	// Name of the transition to make active
	TransitionName string `json:"transitionName,omitempty"`
}

/*
SetCurrentSceneTransitionResponse represents the response body for the "SetCurrentSceneTransition" request.
Sets the current scene transition.

Small note: While the namespace of scene transitions is generally unique, that uniqueness is not a guarantee as it is with other resources like inputs.
*/
type SetCurrentSceneTransitionResponse struct{}

// SetCurrentSceneTransition sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentSceneTransition(
	params *SetCurrentSceneTransitionParams,
) (*SetCurrentSceneTransitionResponse, error) {
	resp, err := c.SendRequest("SetCurrentSceneTransition", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetCurrentSceneTransitionResponse), nil
}
