// This file has been automatically generated. Don't edit it.

package config

import requests "github.com/andreykaipov/goobs/api/requests"

/*
CreateSceneCollectionParams represents the params body for the "CreateSceneCollection" request.
Creates a new scene collection, switching to it in the process.

Note: This will block until the collection has finished changing.
*/
type CreateSceneCollectionParams struct {
	requests.ParamsBasic

	// Name for the new scene collection
	SceneCollectionName string `json:"sceneCollectionName,omitempty"`
}

// GetSelfName just returns "CreateSceneCollection".
func (o *CreateSceneCollectionParams) GetSelfName() string {
	return "CreateSceneCollection"
}

/*
CreateSceneCollectionResponse represents the response body for the "CreateSceneCollection" request.
Creates a new scene collection, switching to it in the process.

Note: This will block until the collection has finished changing.
*/
type CreateSceneCollectionResponse struct {
	requests.ResponseBasic
}

// CreateSceneCollection sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) CreateSceneCollection(params *CreateSceneCollectionParams) (*CreateSceneCollectionResponse, error) {
	data := &CreateSceneCollectionResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
