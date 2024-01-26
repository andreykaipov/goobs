// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the GetSceneItemIndex request.
type GetSceneItemIndexParams struct {
	// Numeric ID of the scene item
	SceneItemId *int `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene the item is in
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewGetSceneItemIndexParams() *GetSceneItemIndexParams {
	return &GetSceneItemIndexParams{}
}
func (o *GetSceneItemIndexParams) WithSceneItemId(x int) *GetSceneItemIndexParams {
	o.SceneItemId = &x
	return o
}
func (o *GetSceneItemIndexParams) WithSceneName(x string) *GetSceneItemIndexParams {
	o.SceneName = &x
	return o
}
func (o *GetSceneItemIndexParams) WithSceneUuid(x string) *GetSceneItemIndexParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *GetSceneItemIndexParams) GetRequestName() string {
	return "GetSceneItemIndex"
}

// Represents the response body for the GetSceneItemIndex request.
type GetSceneItemIndexResponse struct {
	_response

	// Index position of the scene item
	SceneItemIndex int `json:"sceneItemIndex,omitempty"`
}

/*
Gets the index position of a scene item in a scene.

An index of 0 is at the bottom of the source list in the UI.

Scenes and Groups
*/
func (c *Client) GetSceneItemIndex(params *GetSceneItemIndexParams) (*GetSceneItemIndexResponse, error) {
	data := &GetSceneItemIndexResponse{}
	return data, c.client.SendRequest(params, data)
}
