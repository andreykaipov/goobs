// This file has been automatically generated. Don't edit it.

package config

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentSceneCollectionParams represents the params body for the "SetCurrentSceneCollection" request.
Switches to a scene collection.

Note: This will block until the collection has finished changing.
*/
type SetCurrentSceneCollectionParams struct {
	requests.ParamsBasic

	// Name of the scene collection to switch to
	SceneCollectionName string `json:"sceneCollectionName,omitempty"`
}

// GetSelfName just returns "SetCurrentSceneCollection".
func (o *SetCurrentSceneCollectionParams) GetSelfName() string {
	return "SetCurrentSceneCollection"
}

/*
SetCurrentSceneCollectionResponse represents the response body for the "SetCurrentSceneCollection" request.
Switches to a scene collection.

Note: This will block until the collection has finished changing.
*/
type SetCurrentSceneCollectionResponse struct {
	requests.ResponseBasic
}

// SetCurrentSceneCollection sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentSceneCollection(
	params *SetCurrentSceneCollectionParams,
) (*SetCurrentSceneCollectionResponse, error) {
	data := &SetCurrentSceneCollectionResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
