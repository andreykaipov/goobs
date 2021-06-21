// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
CreateSourceParams represents the params body for the "CreateSource" request.
Create a source and add it as a sceneitem to a scene.
Since 4.9.0.
*/
type CreateSourceParams struct {
	requests.ParamsBasic

	// Scene to add the new source to.
	SceneName string `json:"sceneName,omitempty"`

	// Set the created SceneItem as visible or not. Defaults to true
	SetVisible bool `json:"setVisible"`

	// Source kind, Eg. `vlc_source`.
	SourceKind string `json:"sourceKind,omitempty"`

	// Source name.
	SourceName string `json:"sourceName,omitempty"`

	// Source settings data.
	SourceSettings map[string]interface{} `json:"sourceSettings,omitempty"`
}

// GetSelfName just returns "CreateSource".
func (o *CreateSourceParams) GetSelfName() string {
	return "CreateSource"
}

/*
CreateSourceResponse represents the response body for the "CreateSource" request.
Create a source and add it as a sceneitem to a scene.
Since v4.9.0.
*/
type CreateSourceResponse struct {
	requests.ResponseBasic

	// ID of the SceneItem in the scene.
	ItemId int `json:"itemId,omitempty"`
}

// CreateSource sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) CreateSource(params *CreateSourceParams) (*CreateSourceResponse, error) {
	data := &CreateSourceResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
