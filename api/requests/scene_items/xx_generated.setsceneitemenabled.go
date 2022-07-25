// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSceneItemEnabledParams represents the params body for the "SetSceneItemEnabled" request.
Sets the enable state of a scene item.

Scenes and Groups
*/
type SetSceneItemEnabledParams struct {
	requests.ParamsBasic

	// New enable state of the scene item
	SceneItemEnabled bool `json:"sceneItemEnabled,omitempty"`

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "SetSceneItemEnabled".
func (o *SetSceneItemEnabledParams) GetSelfName() string {
	return "SetSceneItemEnabled"
}

/*
SetSceneItemEnabledResponse represents the response body for the "SetSceneItemEnabled" request.
Sets the enable state of a scene item.

Scenes and Groups
*/
type SetSceneItemEnabledResponse struct {
	requests.ResponseBasic
}

// SetSceneItemEnabled sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneItemEnabled(params *SetSceneItemEnabledParams) (*SetSceneItemEnabledResponse, error) {
	data := &SetSceneItemEnabledResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
