// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSceneItemBlendModeParams represents the params body for the "SetSceneItemBlendMode" request.
Sets the blend mode of a scene item.

Scenes and Groups
*/
type SetSceneItemBlendModeParams struct {
	requests.ParamsBasic

	// New blend mode
	SceneItemBlendMode string `json:"sceneItemBlendMode,omitempty"`

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "SetSceneItemBlendMode".
func (o *SetSceneItemBlendModeParams) GetSelfName() string {
	return "SetSceneItemBlendMode"
}

/*
SetSceneItemBlendModeResponse represents the response body for the "SetSceneItemBlendMode" request.
Sets the blend mode of a scene item.

Scenes and Groups
*/
type SetSceneItemBlendModeResponse struct {
	requests.ResponseBasic
}

// SetSceneItemBlendMode sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneItemBlendMode(params *SetSceneItemBlendModeParams) (*SetSceneItemBlendModeResponse, error) {
	data := &SetSceneItemBlendModeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
