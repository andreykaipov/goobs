// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the GetSceneItemLocked request.
type GetSceneItemLockedParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *GetSceneItemLockedParams) GetRequestName() string {
	return "GetSceneItemLocked"
}

// Represents the response body for the GetSceneItemLocked request.
type GetSceneItemLockedResponse struct {
	// Whether the scene item is locked. `true` for locked, `false` for unlocked
	SceneItemLocked bool `json:"sceneItemLocked,omitempty"`
}

/*
Gets the lock state of a scene item.

Scenes and Groups
*/
func (c *Client) GetSceneItemLocked(params *GetSceneItemLockedParams) (*GetSceneItemLockedResponse, error) {
	data := &GetSceneItemLockedResponse{}
	return data, c.SendRequest(params, data)
}
