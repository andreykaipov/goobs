// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the SetSceneItemBlendMode request.
type SetSceneItemBlendModeParams struct {
	// New blend mode
	SceneItemBlendMode string `json:"sceneItemBlendMode,omitempty"`

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *SetSceneItemBlendModeParams) GetRequestName() string {
	return "SetSceneItemBlendMode"
}

// Represents the response body for the SetSceneItemBlendMode request.
type SetSceneItemBlendModeResponse struct{}

/*
Sets the blend mode of a scene item.

Scenes and Groups
*/
func (c *Client) SetSceneItemBlendMode(params *SetSceneItemBlendModeParams) (*SetSceneItemBlendModeResponse, error) {
	data := &SetSceneItemBlendModeResponse{}
	return data, c.SendRequest(params, data)
}
