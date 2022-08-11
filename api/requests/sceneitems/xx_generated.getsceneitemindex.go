// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the GetSceneItemIndex request.
type GetSceneItemIndexParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *GetSceneItemIndexParams) GetRequestName() string {
	return "GetSceneItemIndex"
}

// Represents the response body for the GetSceneItemIndex request.
type GetSceneItemIndexResponse struct {
	// Index position of the scene item
	SceneItemIndex float64 `json:"sceneItemIndex,omitempty"`
}

/*
Gets the index position of a scene item in a scene.

An index of 0 is at the bottom of the source list in the UI.

Scenes and Groups
*/
func (c *Client) GetSceneItemIndex(params *GetSceneItemIndexParams) (*GetSceneItemIndexResponse, error) {
	data := &GetSceneItemIndexResponse{}
	return data, c.SendRequest(params, data)
}
