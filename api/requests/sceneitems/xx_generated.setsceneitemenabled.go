// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the SetSceneItemEnabled request.
type SetSceneItemEnabledParams struct {
	// New enable state of the scene item
	SceneItemEnabled *bool `json:"sceneItemEnabled,omitempty"`

	// Numeric ID of the scene item
	SceneItemId *int `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene the item is in
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewSetSceneItemEnabledParams() *SetSceneItemEnabledParams {
	return &SetSceneItemEnabledParams{}
}
func (o *SetSceneItemEnabledParams) WithSceneItemEnabled(x bool) *SetSceneItemEnabledParams {
	o.SceneItemEnabled = &x
	return o
}
func (o *SetSceneItemEnabledParams) WithSceneItemId(x int) *SetSceneItemEnabledParams {
	o.SceneItemId = &x
	return o
}
func (o *SetSceneItemEnabledParams) WithSceneName(x string) *SetSceneItemEnabledParams {
	o.SceneName = &x
	return o
}
func (o *SetSceneItemEnabledParams) WithSceneUuid(x string) *SetSceneItemEnabledParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *SetSceneItemEnabledParams) GetRequestName() string {
	return "SetSceneItemEnabled"
}

// Represents the response body for the SetSceneItemEnabled request.
type SetSceneItemEnabledResponse struct {
	_response
}

/*
Sets the enable state of a scene item.

Scenes and Groups
*/
func (c *Client) SetSceneItemEnabled(params *SetSceneItemEnabledParams) (*SetSceneItemEnabledResponse, error) {
	data := &SetSceneItemEnabledResponse{}
	return data, c.client.SendRequest(params, data)
}
