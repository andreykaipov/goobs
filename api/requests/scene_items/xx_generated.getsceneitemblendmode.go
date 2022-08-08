// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the GetSceneItemBlendMode request.
type GetSceneItemBlendModeParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *GetSceneItemBlendModeParams) GetRequestName() string {
	return "GetSceneItemBlendMode"
}

// Represents the response body for the GetSceneItemBlendMode request.
type GetSceneItemBlendModeResponse struct {
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
	return data, c.SendRequest(params, data)
}
