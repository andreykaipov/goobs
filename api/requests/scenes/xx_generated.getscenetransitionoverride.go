// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneTransitionOverrideParams represents the params body for the "GetSceneTransitionOverride" request.
Get the current scene transition override.
Since 4.8.0.
*/
type GetSceneTransitionOverrideParams struct {
	requests.ParamsBasic

	// Name of the scene to switch to.
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "GetSceneTransitionOverride".
func (o *GetSceneTransitionOverrideParams) GetSelfName() string {
	return "GetSceneTransitionOverride"
}

/*
GetSceneTransitionOverrideResponse represents the response body for the "GetSceneTransitionOverride" request.
Get the current scene transition override.
Since v4.8.0.
*/
type GetSceneTransitionOverrideResponse struct {
	requests.ResponseBasic

	// Transition duration. `-1` if no override is set.
	TransitionDuration int `json:"transitionDuration,omitempty"`

	// Name of the current overriding transition. Empty string if no override is set.
	TransitionName string `json:"transitionName,omitempty"`
}

// GetSceneTransitionOverride sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneTransitionOverride(
	params *GetSceneTransitionOverrideParams,
) (*GetSceneTransitionOverrideResponse, error) {
	data := &GetSceneTransitionOverrideResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
