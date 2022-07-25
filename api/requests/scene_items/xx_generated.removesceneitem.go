// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
RemoveSceneItemParams represents the params body for the "RemoveSceneItem" request.
Removes a scene item from a scene.

Scenes only
*/
type RemoveSceneItemParams struct {
	requests.ParamsBasic

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "RemoveSceneItem".
func (o *RemoveSceneItemParams) GetSelfName() string {
	return "RemoveSceneItem"
}

/*
RemoveSceneItemResponse represents the response body for the "RemoveSceneItem" request.
Removes a scene item from a scene.

Scenes only
*/
type RemoveSceneItemResponse struct {
	requests.ResponseBasic
}

// RemoveSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RemoveSceneItem(params *RemoveSceneItemParams) (*RemoveSceneItemResponse, error) {
	data := &RemoveSceneItemResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
