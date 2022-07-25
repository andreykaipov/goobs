// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneSceneTransitionOverrideParams represents the params body for the "GetSceneSceneTransitionOverride" request.
Gets the scene transition overridden for a scene.
*/
type GetSceneSceneTransitionOverrideParams struct {
	requests.ParamsBasic

	// Name of the scene
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "GetSceneSceneTransitionOverride".
func (o *GetSceneSceneTransitionOverrideParams) GetSelfName() string {
	return "GetSceneSceneTransitionOverride"
}

/*
GetSceneSceneTransitionOverrideResponse represents the response body for the "GetSceneSceneTransitionOverride" request.
Gets the scene transition overridden for a scene.
*/
type GetSceneSceneTransitionOverrideResponse struct {
	requests.ResponseBasic

	// Duration of the overridden scene transition, else `null`
	TransitionDuration float64 `json:"transitionDuration,omitempty"`

	// Name of the overridden scene transition, else `null`
	TransitionName string `json:"transitionName,omitempty"`
}

// GetSceneSceneTransitionOverride sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneSceneTransitionOverride(
	params *GetSceneSceneTransitionOverrideParams,
) (*GetSceneSceneTransitionOverrideResponse, error) {
	data := &GetSceneSceneTransitionOverrideResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
