// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
RemoveSceneTransitionOverrideParams represents the params body for the "RemoveSceneTransitionOverride" request.
Remove any transition override on a scene.
Since 4.8.0.
*/
type RemoveSceneTransitionOverrideParams struct {
	requests.ParamsBasic

	// Name of the scene to switch to.
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "RemoveSceneTransitionOverride".
func (o *RemoveSceneTransitionOverrideParams) GetSelfName() string {
	return "RemoveSceneTransitionOverride"
}

/*
RemoveSceneTransitionOverrideResponse represents the response body for the "RemoveSceneTransitionOverride" request.
Remove any transition override on a scene.
Since v4.8.0.
*/
type RemoveSceneTransitionOverrideResponse struct {
	requests.ResponseBasic
}

// RemoveSceneTransitionOverride sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RemoveSceneTransitionOverride(
	params *RemoveSceneTransitionOverrideParams,
) (*RemoveSceneTransitionOverrideResponse, error) {
	data := &RemoveSceneTransitionOverrideResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
