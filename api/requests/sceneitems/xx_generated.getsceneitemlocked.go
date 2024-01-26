// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the GetSceneItemLocked request.
type GetSceneItemLockedParams struct {
	// Numeric ID of the scene item
	SceneItemId *int `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene the item is in
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewGetSceneItemLockedParams() *GetSceneItemLockedParams {
	return &GetSceneItemLockedParams{}
}
func (o *GetSceneItemLockedParams) WithSceneItemId(x int) *GetSceneItemLockedParams {
	o.SceneItemId = &x
	return o
}
func (o *GetSceneItemLockedParams) WithSceneName(x string) *GetSceneItemLockedParams {
	o.SceneName = &x
	return o
}
func (o *GetSceneItemLockedParams) WithSceneUuid(x string) *GetSceneItemLockedParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *GetSceneItemLockedParams) GetRequestName() string {
	return "GetSceneItemLocked"
}

// Represents the response body for the GetSceneItemLocked request.
type GetSceneItemLockedResponse struct {
	_response

	// Whether the scene item is locked. `true` for locked, `false` for unlocked
	SceneItemLocked bool `json:"sceneItemLocked,omitempty"`
}

/*
Gets the lock state of a scene item.

Scenes and Groups
*/
func (c *Client) GetSceneItemLocked(params *GetSceneItemLockedParams) (*GetSceneItemLockedResponse, error) {
	data := &GetSceneItemLockedResponse{}
	return data, c.client.SendRequest(params, data)
}
