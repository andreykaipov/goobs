// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
AddSceneItemParams represents the params body for the "AddSceneItem" request.
Creates a scene item in a scene. In other words, this is how you add a source into a scene.
Since 4.9.0.
*/
type AddSceneItemParams struct {
	requests.ParamsBasic

	// Name of the scene to create the scene item in
	SceneName string `json:"sceneName,omitempty"`

	// Whether to make the sceneitem visible on creation or not. Default `true`
	SetVisible bool `json:"setVisible"`

	// Name of the source to be added
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "AddSceneItem".
func (o *AddSceneItemParams) GetSelfName() string {
	return "AddSceneItem"
}

/*
AddSceneItemResponse represents the response body for the "AddSceneItem" request.
Creates a scene item in a scene. In other words, this is how you add a source into a scene.
Since v4.9.0.
*/
type AddSceneItemResponse struct {
	requests.ResponseBasic

	// Numerical ID of the created scene item
	ItemId int `json:"itemId,omitempty"`
}

// AddSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) AddSceneItem(params *AddSceneItemParams) (*AddSceneItemResponse, error) {
	data := &AddSceneItemResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
