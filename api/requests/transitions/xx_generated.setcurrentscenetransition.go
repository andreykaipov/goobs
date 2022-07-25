// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentSceneTransitionParams represents the params body for the "SetCurrentSceneTransition" request.
Sets the current scene transition.

Small note: While the namespace of scene transitions is generally unique, that uniqueness is not a guarantee as it is with other resources like inputs.
*/
type SetCurrentSceneTransitionParams struct {
	requests.ParamsBasic

	// Name of the transition to make active
	TransitionName string `json:"transitionName,omitempty"`
}

// GetSelfName just returns "SetCurrentSceneTransition".
func (o *SetCurrentSceneTransitionParams) GetSelfName() string {
	return "SetCurrentSceneTransition"
}

/*
SetCurrentSceneTransitionResponse represents the response body for the "SetCurrentSceneTransition" request.
Sets the current scene transition.

Small note: While the namespace of scene transitions is generally unique, that uniqueness is not a guarantee as it is with other resources like inputs.
*/
type SetCurrentSceneTransitionResponse struct {
	requests.ResponseBasic
}

// SetCurrentSceneTransition sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentSceneTransition(
	params *SetCurrentSceneTransitionParams,
) (*SetCurrentSceneTransitionResponse, error) {
	data := &SetCurrentSceneTransitionResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
