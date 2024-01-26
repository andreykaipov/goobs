// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the SetSceneItemBlendMode request.
type SetSceneItemBlendModeParams struct {
	// New blend mode
	SceneItemBlendMode *string `json:"sceneItemBlendMode,omitempty"`

	// Numeric ID of the scene item
	SceneItemId *int `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene the item is in
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewSetSceneItemBlendModeParams() *SetSceneItemBlendModeParams {
	return &SetSceneItemBlendModeParams{}
}
func (o *SetSceneItemBlendModeParams) WithSceneItemBlendMode(x string) *SetSceneItemBlendModeParams {
	o.SceneItemBlendMode = &x
	return o
}
func (o *SetSceneItemBlendModeParams) WithSceneItemId(x int) *SetSceneItemBlendModeParams {
	o.SceneItemId = &x
	return o
}
func (o *SetSceneItemBlendModeParams) WithSceneName(x string) *SetSceneItemBlendModeParams {
	o.SceneName = &x
	return o
}
func (o *SetSceneItemBlendModeParams) WithSceneUuid(x string) *SetSceneItemBlendModeParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *SetSceneItemBlendModeParams) GetRequestName() string {
	return "SetSceneItemBlendMode"
}

// Represents the response body for the SetSceneItemBlendMode request.
type SetSceneItemBlendModeResponse struct {
	_response
}

/*
Sets the blend mode of a scene item.

Scenes and Groups
*/
func (c *Client) SetSceneItemBlendMode(params *SetSceneItemBlendModeParams) (*SetSceneItemBlendModeResponse, error) {
	data := &SetSceneItemBlendModeResponse{}
	return data, c.client.SendRequest(params, data)
}
