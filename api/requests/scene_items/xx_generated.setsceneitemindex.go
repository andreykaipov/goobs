// This file has been automatically generated. Don't edit it.

package sceneitems

/*
SetSceneItemIndexParams represents the params body for the "SetSceneItemIndex" request.
Sets the index position of a scene item in a scene.

Scenes and Groups
*/
type SetSceneItemIndexParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// New index position of the scene item
	SceneItemIndex float64 `json:"sceneItemIndex,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

/*
SetSceneItemIndexResponse represents the response body for the "SetSceneItemIndex" request.
Sets the index position of a scene item in a scene.

Scenes and Groups
*/
type SetSceneItemIndexResponse struct{}

// SetSceneItemIndex sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneItemIndex(params *SetSceneItemIndexParams) (*SetSceneItemIndexResponse, error) {
	resp, err := c.SendRequest("SetSceneItemIndex", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetSceneItemIndexResponse), nil
}
