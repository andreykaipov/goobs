// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the DuplicateSceneItem request.
type DuplicateSceneItemParams struct {
	// Name of the scene to create the duplicated item in
	DestinationSceneName *string `json:"destinationSceneName,omitempty"`

	// UUID of the scene to create the duplicated item in
	DestinationSceneUuid *string `json:"destinationSceneUuid,omitempty"`

	// Numeric ID of the scene item
	SceneItemId *int `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene the item is in
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewDuplicateSceneItemParams() *DuplicateSceneItemParams {
	return &DuplicateSceneItemParams{}
}
func (o *DuplicateSceneItemParams) WithDestinationSceneName(x string) *DuplicateSceneItemParams {
	o.DestinationSceneName = &x
	return o
}
func (o *DuplicateSceneItemParams) WithDestinationSceneUuid(x string) *DuplicateSceneItemParams {
	o.DestinationSceneUuid = &x
	return o
}
func (o *DuplicateSceneItemParams) WithSceneItemId(x int) *DuplicateSceneItemParams {
	o.SceneItemId = &x
	return o
}
func (o *DuplicateSceneItemParams) WithSceneName(x string) *DuplicateSceneItemParams {
	o.SceneName = &x
	return o
}
func (o *DuplicateSceneItemParams) WithSceneUuid(x string) *DuplicateSceneItemParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *DuplicateSceneItemParams) GetRequestName() string {
	return "DuplicateSceneItem"
}

// Represents the response body for the DuplicateSceneItem request.
type DuplicateSceneItemResponse struct {
	_response

	// Numeric ID of the duplicated scene item
	SceneItemId int `json:"sceneItemId,omitempty"`
}

/*
Duplicates a scene item, copying all transform and crop info.

Scenes only
*/
func (c *Client) DuplicateSceneItem(params *DuplicateSceneItemParams) (*DuplicateSceneItemResponse, error) {
	data := &DuplicateSceneItemResponse{}
	return data, c.client.SendRequest(params, data)
}
