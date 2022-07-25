// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
CreateSceneItemParams represents the params body for the "CreateSceneItem" request.
Creates a new scene item using a source.

Scenes only
*/
type CreateSceneItemParams struct {
	requests.ParamsBasic

	// Enable state to apply to the scene item on creation
	SceneItemEnabled bool `json:"sceneItemEnabled,omitempty"`

	// Name of the scene to create the new item in
	SceneName string `json:"sceneName,omitempty"`

	// Name of the source to add to the scene
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "CreateSceneItem".
func (o *CreateSceneItemParams) GetSelfName() string {
	return "CreateSceneItem"
}

/*
CreateSceneItemResponse represents the response body for the "CreateSceneItem" request.
Creates a new scene item using a source.

Scenes only
*/
type CreateSceneItemResponse struct {
	requests.ResponseBasic

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`
}

// CreateSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) CreateSceneItem(params *CreateSceneItemParams) (*CreateSceneItemResponse, error) {
	data := &CreateSceneItemResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
