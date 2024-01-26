// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the GetSceneItemBlendMode request.
type GetSceneItemBlendModeParams struct {
	// Numeric ID of the scene item
	SceneItemId *int `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene the item is in
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewGetSceneItemBlendModeParams() *GetSceneItemBlendModeParams {
	return &GetSceneItemBlendModeParams{}
}
func (o *GetSceneItemBlendModeParams) WithSceneItemId(x int) *GetSceneItemBlendModeParams {
	o.SceneItemId = &x
	return o
}
func (o *GetSceneItemBlendModeParams) WithSceneName(x string) *GetSceneItemBlendModeParams {
	o.SceneName = &x
	return o
}
func (o *GetSceneItemBlendModeParams) WithSceneUuid(x string) *GetSceneItemBlendModeParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *GetSceneItemBlendModeParams) GetRequestName() string {
	return "GetSceneItemBlendMode"
}

// Represents the response body for the GetSceneItemBlendMode request.
type GetSceneItemBlendModeResponse struct {
	_response

	// Current blend mode
	SceneItemBlendMode string `json:"sceneItemBlendMode,omitempty"`
}

/*
Gets the blend mode of a scene item.

Blend modes:

- `OBS_BLEND_NORMAL`
- `OBS_BLEND_ADDITIVE`
- `OBS_BLEND_SUBTRACT`
- `OBS_BLEND_SCREEN`
- `OBS_BLEND_MULTIPLY`
- `OBS_BLEND_LIGHTEN`
- `OBS_BLEND_DARKEN`

Scenes and Groups
*/
func (c *Client) GetSceneItemBlendMode(params *GetSceneItemBlendModeParams) (*GetSceneItemBlendModeResponse, error) {
	data := &GetSceneItemBlendModeResponse{}
	return data, c.client.SendRequest(params, data)
}
