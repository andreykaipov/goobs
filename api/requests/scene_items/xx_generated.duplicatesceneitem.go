// This file has been automatically generated. Don't edit it.

package sceneitems

/*
DuplicateSceneItemParams represents the params body for the "DuplicateSceneItem" request.
Duplicates a scene item, copying all transform and crop info.

Scenes only
*/
type DuplicateSceneItemParams struct {
	// Name of the scene to create the duplicated item in
	DestinationSceneName string `json:"destinationSceneName,omitempty"`

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

/*
DuplicateSceneItemResponse represents the response body for the "DuplicateSceneItem" request.
Duplicates a scene item, copying all transform and crop info.

Scenes only
*/
type DuplicateSceneItemResponse struct {
	// Numeric ID of the duplicated scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`
}

// DuplicateSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) DuplicateSceneItem(params *DuplicateSceneItemParams) (*DuplicateSceneItemResponse, error) {
	resp, err := c.SendRequest("DuplicateSceneItem", params)
	if err != nil {
		return nil, err
	}
	return resp.(*DuplicateSceneItemResponse), nil
}
