// This file has been automatically generated. Don't edit it.

package sceneitems

// Represents the request body for the DuplicateSceneItem request.
type DuplicateSceneItemParams struct {
	// Name of the scene to create the duplicated item in
	DestinationSceneName string `json:"destinationSceneName,omitempty"`

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *DuplicateSceneItemParams) GetRequestName() string {
	return "DuplicateSceneItem"
}

// Represents the response body for the DuplicateSceneItem request.
type DuplicateSceneItemResponse struct {
	// Numeric ID of the duplicated scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`
}

/*
Duplicates a scene item, copying all transform and crop info.

Scenes only
*/
func (c *Client) DuplicateSceneItem(params *DuplicateSceneItemParams) (*DuplicateSceneItemResponse, error) {
	data := &DuplicateSceneItemResponse{}
	return data, c.SendRequest(params, data)
}
