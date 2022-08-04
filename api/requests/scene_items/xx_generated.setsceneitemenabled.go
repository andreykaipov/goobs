// This file has been automatically generated. Don't edit it.

package sceneitems

/*
SetSceneItemEnabledParams represents the params body for the "SetSceneItemEnabled" request.
Sets the enable state of a scene item.

Scenes and Groups
*/
type SetSceneItemEnabledParams struct {
	// New enable state of the scene item
	SceneItemEnabled *bool `json:"sceneItemEnabled,omitempty"`

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

/*
SetSceneItemEnabledResponse represents the response body for the "SetSceneItemEnabled" request.
Sets the enable state of a scene item.

Scenes and Groups
*/
type SetSceneItemEnabledResponse struct{}

// SetSceneItemEnabled sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneItemEnabled(params *SetSceneItemEnabledParams) (*SetSceneItemEnabledResponse, error) {
	resp, err := c.SendRequest("SetSceneItemEnabled", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetSceneItemEnabledResponse), nil
}
