// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentSceneParams represents the params body for the "SetCurrentScene" request.
Switch to the specified scene.
Since 0.3.
*/
type SetCurrentSceneParams struct {
	requests.ParamsBasic

	// Name of the scene to switch to.
	SceneName string `json:"scene-name,omitempty"`
}

// GetSelfName just returns "SetCurrentScene".
func (o *SetCurrentSceneParams) GetSelfName() string {
	return "SetCurrentScene"
}

/*
SetCurrentSceneResponse represents the response body for the "SetCurrentScene" request.
Switch to the specified scene.
Since v0.3.
*/
type SetCurrentSceneResponse struct {
	requests.ResponseBasic
}

// SetCurrentScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentScene(params *SetCurrentSceneParams) (*SetCurrentSceneResponse, error) {
	data := &SetCurrentSceneResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
