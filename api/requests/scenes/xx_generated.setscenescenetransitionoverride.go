// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSceneSceneTransitionOverrideParams represents the params body for the "SetSceneSceneTransitionOverride" request.
Gets the scene transition overridden for a scene.
*/
type SetSceneSceneTransitionOverrideParams struct {
	requests.ParamsBasic

	// Name of the scene
	SceneName string `json:"sceneName,omitempty"`

	// Duration to use for any overridden transition. Specify `null` to remove
	TransitionDuration float64 `json:"transitionDuration,omitempty"`

	// Name of the scene transition to use as override. Specify `null` to remove
	TransitionName string `json:"transitionName,omitempty"`
}

// GetSelfName just returns "SetSceneSceneTransitionOverride".
func (o *SetSceneSceneTransitionOverrideParams) GetSelfName() string {
	return "SetSceneSceneTransitionOverride"
}

/*
SetSceneSceneTransitionOverrideResponse represents the response body for the "SetSceneSceneTransitionOverride" request.
Gets the scene transition overridden for a scene.
*/
type SetSceneSceneTransitionOverrideResponse struct {
	requests.ResponseBasic
}

// SetSceneSceneTransitionOverride sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneSceneTransitionOverride(
	params *SetSceneSceneTransitionOverrideParams,
) (*SetSceneSceneTransitionOverrideResponse, error) {
	data := &SetSceneSceneTransitionOverrideResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
