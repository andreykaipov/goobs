// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the SetSceneItemIndex request.
type SetSceneItemIndexParams struct {
	// Numeric ID of the scene item
	SceneItemId *int `json:"sceneItemId,omitempty"`

	// New index position of the scene item
	SceneItemIndex *int `json:"sceneItemIndex,omitempty"`

	// Name of the scene the item is in
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene the item is in
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewSetSceneItemIndexParams() *SetSceneItemIndexParams {
	return &SetSceneItemIndexParams{}
}
func (o *SetSceneItemIndexParams) WithSceneItemId(x int) *SetSceneItemIndexParams {
	o.SceneItemId = &x
	return o
}
func (o *SetSceneItemIndexParams) WithSceneItemIndex(x int) *SetSceneItemIndexParams {
	o.SceneItemIndex = &x
	return o
}
func (o *SetSceneItemIndexParams) WithSceneName(x string) *SetSceneItemIndexParams {
	o.SceneName = &x
	return o
}
func (o *SetSceneItemIndexParams) WithSceneUuid(x string) *SetSceneItemIndexParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *SetSceneItemIndexParams) GetRequestName() string {
	return "SetSceneItemIndex"
}

// Represents the response body for the SetSceneItemIndex request.
type SetSceneItemIndexResponse struct {
	_response
}

/*
Sets the index position of a scene item in a scene.

Scenes and Groups
*/
func (c *Client) SetSceneItemIndex(params *SetSceneItemIndexParams) (*SetSceneItemIndexResponse, error) {
	data := &SetSceneItemIndexResponse{}
	return data, c.client.SendRequest(params, data)
}
