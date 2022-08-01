// This file has been automatically generated. Don't edit it.

package sceneitems

/*
GetSceneItemLockedParams represents the params body for the "GetSceneItemLocked" request.
Gets the lock state of a scene item.

Scenes and Groups
*/
type GetSceneItemLockedParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

/*
GetSceneItemLockedResponse represents the response body for the "GetSceneItemLocked" request.
Gets the lock state of a scene item.

Scenes and Groups
*/
type GetSceneItemLockedResponse struct {
	// Whether the scene item is locked. `true` for locked, `false` for unlocked
	SceneItemLocked bool `json:"sceneItemLocked,omitempty"`
}

// GetSceneItemLocked sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemLocked(params *GetSceneItemLockedParams) (*GetSceneItemLockedResponse, error) {
	resp, err := c.SendRequest("GetSceneItemLocked", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetSceneItemLockedResponse), nil
}
