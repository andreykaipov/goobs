// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
DuplicateSceneItemParams represents the params body for the "DuplicateSceneItem" request.
Duplicates a scene item, copying all transform and crop info.

Scenes only
*/
type DuplicateSceneItemParams struct {
	requests.ParamsBasic

	// Name of the scene to create the duplicated item in
	DestinationSceneName string `json:"destinationSceneName,omitempty"`

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "DuplicateSceneItem".
func (o *DuplicateSceneItemParams) GetSelfName() string {
	return "DuplicateSceneItem"
}

/*
DuplicateSceneItemResponse represents the response body for the "DuplicateSceneItem" request.
Duplicates a scene item, copying all transform and crop info.

Scenes only
*/
type DuplicateSceneItemResponse struct {
	requests.ResponseBasic

	// Numeric ID of the duplicated scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`
}

// DuplicateSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) DuplicateSceneItem(params *DuplicateSceneItemParams) (*DuplicateSceneItemResponse, error) {
	data := &DuplicateSceneItemResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
