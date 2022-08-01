// This file has been automatically generated. Don't edit it.

package sceneitems

/*
GetSceneItemEnabledParams represents the params body for the "GetSceneItemEnabled" request.
Gets the enable state of a scene item.

Scenes and Groups
*/
type GetSceneItemEnabledParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

/*
GetSceneItemEnabledResponse represents the response body for the "GetSceneItemEnabled" request.
Gets the enable state of a scene item.

Scenes and Groups
*/
type GetSceneItemEnabledResponse struct {
	// Whether the scene item is enabled. `true` for enabled, `false` for disabled
	SceneItemEnabled bool `json:"sceneItemEnabled,omitempty"`
}

// GetSceneItemEnabled sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemEnabled(params *GetSceneItemEnabledParams) (*GetSceneItemEnabledResponse, error) {
	resp, err := c.SendRequest("GetSceneItemEnabled", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetSceneItemEnabledResponse), nil
}
