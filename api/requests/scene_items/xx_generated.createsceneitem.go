// This file has been automatically generated. Don't edit it.

package sceneitems

/*
CreateSceneItemParams represents the params body for the "CreateSceneItem" request.
Creates a new scene item using a source.

Scenes only
*/
type CreateSceneItemParams struct {
	// Enable state to apply to the scene item on creation
	SceneItemEnabled bool `json:"sceneItemEnabled,omitempty"`

	// Name of the scene to create the new item in
	SceneName string `json:"sceneName,omitempty"`

	// Name of the source to add to the scene
	SourceName string `json:"sourceName,omitempty"`
}

/*
CreateSceneItemResponse represents the response body for the "CreateSceneItem" request.
Creates a new scene item using a source.

Scenes only
*/
type CreateSceneItemResponse struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`
}

// CreateSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) CreateSceneItem(params *CreateSceneItemParams) (*CreateSceneItemResponse, error) {
	resp, err := c.SendRequest("CreateSceneItem", params)
	if err != nil {
		return nil, err
	}
	return resp.(*CreateSceneItemResponse), nil
}
