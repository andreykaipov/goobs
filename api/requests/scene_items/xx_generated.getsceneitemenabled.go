// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneItemEnabledParams represents the params body for the "GetSceneItemEnabled" request.
Gets the enable state of a scene item.

Scenes and Groups
*/
type GetSceneItemEnabledParams struct {
	requests.ParamsBasic

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "GetSceneItemEnabled".
func (o *GetSceneItemEnabledParams) GetSelfName() string {
	return "GetSceneItemEnabled"
}

/*
GetSceneItemEnabledResponse represents the response body for the "GetSceneItemEnabled" request.
Gets the enable state of a scene item.

Scenes and Groups
*/
type GetSceneItemEnabledResponse struct {
	requests.ResponseBasic

	// Whether the scene item is enabled. `true` for enabled, `false` for disabled
	SceneItemEnabled bool `json:"sceneItemEnabled,omitempty"`
}

// GetSceneItemEnabled sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemEnabled(params *GetSceneItemEnabledParams) (*GetSceneItemEnabledResponse, error) {
	data := &GetSceneItemEnabledResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
