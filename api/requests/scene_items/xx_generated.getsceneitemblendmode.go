// This file has been automatically generated. Don't edit it.

package sceneitems

/*
GetSceneItemBlendModeParams represents the params body for the "GetSceneItemBlendMode" request.
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
type GetSceneItemBlendModeParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

/*
GetSceneItemBlendModeResponse represents the response body for the "GetSceneItemBlendMode" request.
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
type GetSceneItemBlendModeResponse struct {
	// Current blend mode
	SceneItemBlendMode string `json:"sceneItemBlendMode,omitempty"`
}

// GetSceneItemBlendMode sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemBlendMode(params *GetSceneItemBlendModeParams) (*GetSceneItemBlendModeResponse, error) {
	resp, err := c.SendRequest("GetSceneItemBlendMode", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetSceneItemBlendModeResponse), nil
}
