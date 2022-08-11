// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the SetSceneItemIndex request.
type SetSceneItemIndexParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// New index position of the scene item
	SceneItemIndex float64 `json:"sceneItemIndex,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *SetSceneItemIndexParams) GetRequestName() string {
	return "SetSceneItemIndex"
}

// Represents the response body for the SetSceneItemIndex request.
type SetSceneItemIndexResponse struct{}

/*
Sets the index position of a scene item in a scene.

Scenes and Groups
*/
func (c *Client) SetSceneItemIndex(params *SetSceneItemIndexParams) (*SetSceneItemIndexResponse, error) {
	data := &SetSceneItemIndexResponse{}
	return data, c.SendRequest(params, data)
}
