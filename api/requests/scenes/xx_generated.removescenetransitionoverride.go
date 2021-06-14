// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
RemoveSceneTransitionOverrideParams represents the params body for the "RemoveSceneTransitionOverride" request.
Remove any transition override on a scene.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#RemoveSceneTransitionOverride.
*/
type RemoveSceneTransitionOverrideParams struct {
	requests.ParamsBasic

	// Name of the scene to switch to.
	SceneName string `json:"sceneName"`
}

// GetSelfName just returns "RemoveSceneTransitionOverride".
func (o *RemoveSceneTransitionOverrideParams) GetSelfName() string {
	return "RemoveSceneTransitionOverride"
}

/*
RemoveSceneTransitionOverrideResponse represents the response body for the "RemoveSceneTransitionOverride" request.
Remove any transition override on a scene.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#RemoveSceneTransitionOverride.
*/
type RemoveSceneTransitionOverrideResponse struct {
	requests.ResponseBasic
}

// RemoveSceneTransitionOverride sends the corresponding request to the connected OBS WebSockets
// server.
func (c *Client) RemoveSceneTransitionOverride(
	params *RemoveSceneTransitionOverrideParams,
) (*RemoveSceneTransitionOverrideResponse, error) {
	data := &RemoveSceneTransitionOverrideResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
