// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentSceneParams represents the params body for the "SetCurrentScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentScene.
*/
type SetCurrentSceneParams struct {
	requests.ParamsBasic

	// Name of the scene to switch to.
	SceneName string `json:"scene-name"`
}

// Name just returns "SetCurrentScene".
func (o *SetCurrentSceneParams) Name() string {
	return "SetCurrentScene"
}

/*
SetCurrentSceneResponse represents the response body for the "SetCurrentScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentScene.
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
