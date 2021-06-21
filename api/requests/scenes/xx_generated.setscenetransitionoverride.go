// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSceneTransitionOverrideParams represents the params body for the "SetSceneTransitionOverride" request.
Set a scene to use a specific transition override.
Since 4.8.0.
*/
type SetSceneTransitionOverrideParams struct {
	requests.ParamsBasic

	// Name of the scene to switch to.
	SceneName string `json:"sceneName,omitempty"`

	// Duration in milliseconds of the transition if transition is not fixed. Defaults to the current duration specified
	// in the UI if there is no current override and this value is not given.
	TransitionDuration int `json:"transitionDuration,omitempty"`

	// Name of the transition to use.
	TransitionName string `json:"transitionName,omitempty"`
}

// GetSelfName just returns "SetSceneTransitionOverride".
func (o *SetSceneTransitionOverrideParams) GetSelfName() string {
	return "SetSceneTransitionOverride"
}

/*
SetSceneTransitionOverrideResponse represents the response body for the "SetSceneTransitionOverride" request.
Set a scene to use a specific transition override.
Since v4.8.0.
*/
type SetSceneTransitionOverrideResponse struct {
	requests.ResponseBasic
}

// SetSceneTransitionOverride sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneTransitionOverride(
	params *SetSceneTransitionOverrideParams,
) (*SetSceneTransitionOverrideResponse, error) {
	data := &SetSceneTransitionOverrideResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
