// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

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
	requests.ParamsBasic

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "GetSceneItemBlendMode".
func (o *GetSceneItemBlendModeParams) GetSelfName() string {
	return "GetSceneItemBlendMode"
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
	requests.ResponseBasic

	// Current blend mode
	SceneItemBlendMode string `json:"sceneItemBlendMode,omitempty"`
}

// GetSceneItemBlendMode sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemBlendMode(params *GetSceneItemBlendModeParams) (*GetSceneItemBlendModeResponse, error) {
	data := &GetSceneItemBlendModeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
