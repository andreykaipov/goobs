// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the SetCurrentSceneCollection request.
type SetCurrentSceneCollectionParams struct {
	// Name of the scene collection to switch to
	SceneCollectionName string `json:"sceneCollectionName,omitempty"`
}

// Returns the associated request.
func (o *SetCurrentSceneCollectionParams) GetRequestName() string {
	return "SetCurrentSceneCollection"
}

// Represents the response body for the SetCurrentSceneCollection request.
type SetCurrentSceneCollectionResponse struct{}

/*
Switches to a scene collection.

Note: This will block until the collection has finished changing.
*/
func (c *Client) SetCurrentSceneCollection(
	params *SetCurrentSceneCollectionParams,
) (*SetCurrentSceneCollectionResponse, error) {
	data := &SetCurrentSceneCollectionResponse{}
	return data, c.SendRequest(params, data)
}
