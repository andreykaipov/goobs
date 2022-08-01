// This file has been automatically generated. Don't edit it.

package sceneitems

/*
RemoveSceneItemParams represents the params body for the "RemoveSceneItem" request.
Removes a scene item from a scene.

Scenes only
*/
type RemoveSceneItemParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

/*
RemoveSceneItemResponse represents the response body for the "RemoveSceneItem" request.
Removes a scene item from a scene.

Scenes only
*/
type RemoveSceneItemResponse struct{}

// RemoveSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RemoveSceneItem(params *RemoveSceneItemParams) (*RemoveSceneItemResponse, error) {
	resp, err := c.SendRequest("RemoveSceneItem", params)
	if err != nil {
		return nil, err
	}
	return resp.(*RemoveSceneItemResponse), nil
}
