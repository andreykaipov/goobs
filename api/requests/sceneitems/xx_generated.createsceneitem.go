// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the CreateSceneItem request.
type CreateSceneItemParams struct {
	// Enable state to apply to the scene item on creation
	SceneItemEnabled *bool `json:"sceneItemEnabled,omitempty"`

	// Name of the scene to create the new item in
	SceneName *string `json:"sceneName,omitempty"`

	// Name of the source to add to the scene
	SourceName *string `json:"sourceName,omitempty"`
}

func NewCreateSceneItemParams() *CreateSceneItemParams {
	return &CreateSceneItemParams{}
}
func (o *CreateSceneItemParams) WithSceneItemEnabled(x bool) *CreateSceneItemParams {
	o.SceneItemEnabled = &x
	return o
}
func (o *CreateSceneItemParams) WithSceneName(x string) *CreateSceneItemParams {
	o.SceneName = &x
	return o
}
func (o *CreateSceneItemParams) WithSourceName(x string) *CreateSceneItemParams {
	o.SourceName = &x
	return o
}

// Returns the associated request.
func (o *CreateSceneItemParams) GetRequestName() string {
	return "CreateSceneItem"
}

// Represents the response body for the CreateSceneItem request.
type CreateSceneItemResponse struct {
	_response

	// Numeric ID of the scene item
	SceneItemId int `json:"sceneItemId,omitempty"`
}

/*
Creates a new scene item using a source.

Scenes only
*/
func (c *Client) CreateSceneItem(params *CreateSceneItemParams) (*CreateSceneItemResponse, error) {
	data := &CreateSceneItemResponse{}
	return data, c.SendRequest(params, data)
}
