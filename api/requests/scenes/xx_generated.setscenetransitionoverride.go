// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSceneTransitionOverrideParams represents the params body for the "SetSceneTransitionOverride" request.
Set a scene to use a specific transition override.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SetSceneTransitionOverride.
*/
type SetSceneTransitionOverrideParams struct {
	requests.ParamsBasic

	// Name of the scene to switch to.
	SceneName string `json:"sceneName"`

	// Duration in milliseconds of the transition if transition is not fixed. Defaults to the
	// current duration specified in the UI if there is no current override and this value is not
	// given.
	TransitionDuration int `json:"transitionDuration"`

	// Name of the transition to use.
	TransitionName string `json:"transitionName"`
}

// Name just returns "SetSceneTransitionOverride".
func (o *SetSceneTransitionOverrideParams) Name() string {
	return "SetSceneTransitionOverride"
}

/*
SetSceneTransitionOverrideResponse represents the response body for the "SetSceneTransitionOverride" request.
Set a scene to use a specific transition override.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SetSceneTransitionOverride.
*/
type SetSceneTransitionOverrideResponse struct {
	requests.ResponseBasic
}

// SetSceneTransitionOverride sends the corresponding request to the connected OBS WebSockets
// server.
func (c *Client) SetSceneTransitionOverride(
	params *SetSceneTransitionOverrideParams,
) (*SetSceneTransitionOverrideResponse, error) {
	data := &SetSceneTransitionOverrideResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
