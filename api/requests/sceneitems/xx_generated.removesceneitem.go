// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the RemoveSceneItem request.
type RemoveSceneItemParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *RemoveSceneItemParams) GetRequestName() string {
	return "RemoveSceneItem"
}

// Represents the response body for the RemoveSceneItem request.
type RemoveSceneItemResponse struct{}

/*
Removes a scene item from a scene.

Scenes only
*/
func (c *Client) RemoveSceneItem(params *RemoveSceneItemParams) (*RemoveSceneItemResponse, error) {
	data := &RemoveSceneItemResponse{}
	return data, c.SendRequest(params, data)
}
