// This file has been automatically generated. Don't edit it.

package sceneitems

/*
GetSceneItemIndexParams represents the params body for the "GetSceneItemIndex" request.
Gets the index position of a scene item in a scene.

An index of 0 is at the bottom of the source list in the UI.

Scenes and Groups
*/
type GetSceneItemIndexParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

/*
GetSceneItemIndexResponse represents the response body for the "GetSceneItemIndex" request.
Gets the index position of a scene item in a scene.

An index of 0 is at the bottom of the source list in the UI.

Scenes and Groups
*/
type GetSceneItemIndexResponse struct {
	// Index position of the scene item
	SceneItemIndex float64 `json:"sceneItemIndex,omitempty"`
}

// GetSceneItemIndex sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemIndex(params *GetSceneItemIndexParams) (*GetSceneItemIndexResponse, error) {
	resp, err := c.SendRequest("GetSceneItemIndex", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetSceneItemIndexResponse), nil
}
