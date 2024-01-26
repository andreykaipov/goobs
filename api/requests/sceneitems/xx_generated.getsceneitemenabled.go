// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the GetSceneItemEnabled request.
type GetSceneItemEnabledParams struct {
	// Numeric ID of the scene item
	SceneItemId *int `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene the item is in
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewGetSceneItemEnabledParams() *GetSceneItemEnabledParams {
	return &GetSceneItemEnabledParams{}
}
func (o *GetSceneItemEnabledParams) WithSceneItemId(x int) *GetSceneItemEnabledParams {
	o.SceneItemId = &x
	return o
}
func (o *GetSceneItemEnabledParams) WithSceneName(x string) *GetSceneItemEnabledParams {
	o.SceneName = &x
	return o
}
func (o *GetSceneItemEnabledParams) WithSceneUuid(x string) *GetSceneItemEnabledParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *GetSceneItemEnabledParams) GetRequestName() string {
	return "GetSceneItemEnabled"
}

// Represents the response body for the GetSceneItemEnabled request.
type GetSceneItemEnabledResponse struct {
	_response

	// Whether the scene item is enabled. `true` for enabled, `false` for disabled
	SceneItemEnabled bool `json:"sceneItemEnabled,omitempty"`
}

/*
Gets the enable state of a scene item.

Scenes and Groups
*/
func (c *Client) GetSceneItemEnabled(params *GetSceneItemEnabledParams) (*GetSceneItemEnabledResponse, error) {
	data := &GetSceneItemEnabledResponse{}
	return data, c.client.SendRequest(params, data)
}
