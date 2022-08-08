// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the SetSceneItemEnabled request.
type SetSceneItemEnabledParams struct {
	// New enable state of the scene item
	SceneItemEnabled *bool `json:"sceneItemEnabled,omitempty"`

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *SetSceneItemEnabledParams) GetRequestName() string {
	return "SetSceneItemEnabled"
}

// Represents the response body for the SetSceneItemEnabled request.
type SetSceneItemEnabledResponse struct{}

/*
Sets the enable state of a scene item.

Scenes and Groups
*/
func (c *Client) SetSceneItemEnabled(params *SetSceneItemEnabledParams) (*SetSceneItemEnabledResponse, error) {
	data := &SetSceneItemEnabledResponse{}
	return data, c.SendRequest(params, data)
}
