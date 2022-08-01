// This file has been automatically generated. Don't edit it.

package sceneitems

/*
SetSceneItemLockedParams represents the params body for the "SetSceneItemLocked" request.
Sets the lock state of a scene item.

Scenes and Group
*/
type SetSceneItemLockedParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// New lock state of the scene item
	SceneItemLocked bool `json:"sceneItemLocked,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

/*
SetSceneItemLockedResponse represents the response body for the "SetSceneItemLocked" request.
Sets the lock state of a scene item.

Scenes and Group
*/
type SetSceneItemLockedResponse struct{}

// SetSceneItemLocked sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneItemLocked(params *SetSceneItemLockedParams) (*SetSceneItemLockedResponse, error) {
	resp, err := c.SendRequest("SetSceneItemLocked", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetSceneItemLockedResponse), nil
}
