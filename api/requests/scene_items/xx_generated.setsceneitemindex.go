// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSceneItemIndexParams represents the params body for the "SetSceneItemIndex" request.
Sets the index position of a scene item in a scene.

Scenes and Groups
*/
type SetSceneItemIndexParams struct {
	requests.ParamsBasic

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// New index position of the scene item
	SceneItemIndex float64 `json:"sceneItemIndex,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "SetSceneItemIndex".
func (o *SetSceneItemIndexParams) GetSelfName() string {
	return "SetSceneItemIndex"
}

/*
SetSceneItemIndexResponse represents the response body for the "SetSceneItemIndex" request.
Sets the index position of a scene item in a scene.

Scenes and Groups
*/
type SetSceneItemIndexResponse struct {
	requests.ResponseBasic
}

// SetSceneItemIndex sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneItemIndex(params *SetSceneItemIndexParams) (*SetSceneItemIndexResponse, error) {
	data := &SetSceneItemIndexResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
