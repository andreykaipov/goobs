// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the GetSceneItemSource request.
type GetSceneItemSourceParams struct {
	// Numeric ID of the scene item
	SceneItemId *int `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene the item is in
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewGetSceneItemSourceParams() *GetSceneItemSourceParams {
	return &GetSceneItemSourceParams{}
}
func (o *GetSceneItemSourceParams) WithSceneItemId(x int) *GetSceneItemSourceParams {
	o.SceneItemId = &x
	return o
}
func (o *GetSceneItemSourceParams) WithSceneName(x string) *GetSceneItemSourceParams {
	o.SceneName = &x
	return o
}
func (o *GetSceneItemSourceParams) WithSceneUuid(x string) *GetSceneItemSourceParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *GetSceneItemSourceParams) GetRequestName() string {
	return "GetSceneItemSource"
}

// Represents the response body for the GetSceneItemSource request.
type GetSceneItemSourceResponse struct {
	_response

	// Name of the source associated with the scene item
	SourceName string `json:"sourceName,omitempty"`

	// UUID of the source associated with the scene item
	SourceUuid string `json:"sourceUuid,omitempty"`
}

// Gets the source associated with a scene item.
func (c *Client) GetSceneItemSource(params *GetSceneItemSourceParams) (*GetSceneItemSourceResponse, error) {
	data := &GetSceneItemSourceResponse{}
	return data, c.client.SendRequest(params, data)
}
