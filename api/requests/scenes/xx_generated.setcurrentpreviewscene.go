// This file has been automatically generated. Don't edit it.

package scenes

/*
SetCurrentPreviewSceneParams represents the params body for the "SetCurrentPreviewScene" request.
Sets the current preview scene.

Only available when studio mode is enabled.
*/
type SetCurrentPreviewSceneParams struct {
	// Scene to set as the current preview scene
	SceneName string `json:"sceneName,omitempty"`
}

/*
SetCurrentPreviewSceneResponse represents the response body for the "SetCurrentPreviewScene" request.
Sets the current preview scene.

Only available when studio mode is enabled.
*/
type SetCurrentPreviewSceneResponse struct{}

// SetCurrentPreviewScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentPreviewScene(params *SetCurrentPreviewSceneParams) (*SetCurrentPreviewSceneResponse, error) {
	resp, err := c.SendRequest("SetCurrentPreviewScene", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetCurrentPreviewSceneResponse), nil
}
