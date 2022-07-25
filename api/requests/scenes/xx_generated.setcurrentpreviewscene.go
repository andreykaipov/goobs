// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentPreviewSceneParams represents the params body for the "SetCurrentPreviewScene" request.
Sets the current preview scene.

Only available when studio mode is enabled.
*/
type SetCurrentPreviewSceneParams struct {
	requests.ParamsBasic

	// Scene to set as the current preview scene
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "SetCurrentPreviewScene".
func (o *SetCurrentPreviewSceneParams) GetSelfName() string {
	return "SetCurrentPreviewScene"
}

/*
SetCurrentPreviewSceneResponse represents the response body for the "SetCurrentPreviewScene" request.
Sets the current preview scene.

Only available when studio mode is enabled.
*/
type SetCurrentPreviewSceneResponse struct {
	requests.ResponseBasic
}

// SetCurrentPreviewScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentPreviewScene(params *SetCurrentPreviewSceneParams) (*SetCurrentPreviewSceneResponse, error) {
	data := &SetCurrentPreviewSceneResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
