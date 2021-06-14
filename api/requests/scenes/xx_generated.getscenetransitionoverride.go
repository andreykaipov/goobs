// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneTransitionOverrideParams represents the params body for the "GetSceneTransitionOverride" request.
Get the current scene transition override.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetSceneTransitionOverride.
*/
type GetSceneTransitionOverrideParams struct {
	requests.ParamsBasic

	// Name of the scene to switch to.
	SceneName string `json:"sceneName"`
}

// Name just returns "GetSceneTransitionOverride".
func (o *GetSceneTransitionOverrideParams) Name() string {
	return "GetSceneTransitionOverride"
}

/*
GetSceneTransitionOverrideResponse represents the response body for the "GetSceneTransitionOverride" request.
Get the current scene transition override.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetSceneTransitionOverride.
*/
type GetSceneTransitionOverrideResponse struct {
	requests.ResponseBasic

	// Transition duration. `-1` if no override is set.
	TransitionDuration int `json:"transitionDuration"`

	// Name of the current overriding transition. Empty string if no override is set.
	TransitionName string `json:"transitionName"`
}

// GetSceneTransitionOverride sends the corresponding request to the connected OBS WebSockets
// server.
func (c *Client) GetSceneTransitionOverride(
	params *GetSceneTransitionOverrideParams,
) (*GetSceneTransitionOverrideResponse, error) {
	data := &GetSceneTransitionOverrideResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
