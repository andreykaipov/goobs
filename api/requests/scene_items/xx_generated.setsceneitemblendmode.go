// This file has been automatically generated. Don't edit it.

package sceneitems

/*
SetSceneItemBlendModeParams represents the params body for the "SetSceneItemBlendMode" request.
Sets the blend mode of a scene item.

Scenes and Groups
*/
type SetSceneItemBlendModeParams struct {
	// New blend mode
	SceneItemBlendMode string `json:"sceneItemBlendMode,omitempty"`

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

/*
SetSceneItemBlendModeResponse represents the response body for the "SetSceneItemBlendMode" request.
Sets the blend mode of a scene item.

Scenes and Groups
*/
type SetSceneItemBlendModeResponse struct{}

// SetSceneItemBlendMode sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneItemBlendMode(params *SetSceneItemBlendModeParams) (*SetSceneItemBlendModeResponse, error) {
	resp, err := c.SendRequest("SetSceneItemBlendMode", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetSceneItemBlendModeResponse), nil
}
