// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the RemoveSceneItem request.
type RemoveSceneItemParams struct {
	// Numeric ID of the scene item
	SceneItemId *int `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene the item is in
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewRemoveSceneItemParams() *RemoveSceneItemParams {
	return &RemoveSceneItemParams{}
}
func (o *RemoveSceneItemParams) WithSceneItemId(x int) *RemoveSceneItemParams {
	o.SceneItemId = &x
	return o
}
func (o *RemoveSceneItemParams) WithSceneName(x string) *RemoveSceneItemParams {
	o.SceneName = &x
	return o
}
func (o *RemoveSceneItemParams) WithSceneUuid(x string) *RemoveSceneItemParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *RemoveSceneItemParams) GetRequestName() string {
	return "RemoveSceneItem"
}

// Represents the response body for the RemoveSceneItem request.
type RemoveSceneItemResponse struct {
	_response
}

/*
Removes a scene item from a scene.

Scenes only
*/
func (c *Client) RemoveSceneItem(params *RemoveSceneItemParams) (*RemoveSceneItemResponse, error) {
	data := &RemoveSceneItemResponse{}
	return data, c.client.SendRequest(params, data)
}
