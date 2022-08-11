// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the SetSceneItemLocked request.
type SetSceneItemLockedParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// New lock state of the scene item
	SceneItemLocked *bool `json:"sceneItemLocked,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *SetSceneItemLockedParams) GetRequestName() string {
	return "SetSceneItemLocked"
}

// Represents the response body for the SetSceneItemLocked request.
type SetSceneItemLockedResponse struct{}

/*
Sets the lock state of a scene item.

Scenes and Group
*/
func (c *Client) SetSceneItemLocked(params *SetSceneItemLockedParams) (*SetSceneItemLockedResponse, error) {
	data := &SetSceneItemLockedResponse{}
	return data, c.SendRequest(params, data)
}
