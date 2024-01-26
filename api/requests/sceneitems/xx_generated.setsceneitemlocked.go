// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the SetSceneItemLocked request.
type SetSceneItemLockedParams struct {
	// Numeric ID of the scene item
	SceneItemId *int `json:"sceneItemId,omitempty"`

	// New lock state of the scene item
	SceneItemLocked *bool `json:"sceneItemLocked,omitempty"`

	// Name of the scene the item is in
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene the item is in
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewSetSceneItemLockedParams() *SetSceneItemLockedParams {
	return &SetSceneItemLockedParams{}
}
func (o *SetSceneItemLockedParams) WithSceneItemId(x int) *SetSceneItemLockedParams {
	o.SceneItemId = &x
	return o
}
func (o *SetSceneItemLockedParams) WithSceneItemLocked(x bool) *SetSceneItemLockedParams {
	o.SceneItemLocked = &x
	return o
}
func (o *SetSceneItemLockedParams) WithSceneName(x string) *SetSceneItemLockedParams {
	o.SceneName = &x
	return o
}
func (o *SetSceneItemLockedParams) WithSceneUuid(x string) *SetSceneItemLockedParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *SetSceneItemLockedParams) GetRequestName() string {
	return "SetSceneItemLocked"
}

// Represents the response body for the SetSceneItemLocked request.
type SetSceneItemLockedResponse struct {
	_response
}

/*
Sets the lock state of a scene item.

Scenes and Group
*/
func (c *Client) SetSceneItemLocked(params *SetSceneItemLockedParams) (*SetSceneItemLockedResponse, error) {
	data := &SetSceneItemLockedResponse{}
	return data, c.client.SendRequest(params, data)
}
