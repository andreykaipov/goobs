// This file has been automatically generated. Don't edit it.

package config

/*
CreateSceneCollectionParams represents the params body for the "CreateSceneCollection" request.
Creates a new scene collection, switching to it in the process.

Note: This will block until the collection has finished changing.
*/
type CreateSceneCollectionParams struct {
	// Name for the new scene collection
	SceneCollectionName string `json:"sceneCollectionName,omitempty"`
}

/*
CreateSceneCollectionResponse represents the response body for the "CreateSceneCollection" request.
Creates a new scene collection, switching to it in the process.

Note: This will block until the collection has finished changing.
*/
type CreateSceneCollectionResponse struct{}

// CreateSceneCollection sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) CreateSceneCollection(params *CreateSceneCollectionParams) (*CreateSceneCollectionResponse, error) {
	resp, err := c.SendRequest("CreateSceneCollection", params)
	if err != nil {
		return nil, err
	}
	return resp.(*CreateSceneCollectionResponse), nil
}
