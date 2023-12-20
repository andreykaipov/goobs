// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the CreateSceneCollection request.
type CreateSceneCollectionParams struct {
	// Name for the new scene collection
	SceneCollectionName *string `json:"sceneCollectionName,omitempty"`
}

func NewCreateSceneCollectionParams() *CreateSceneCollectionParams {
	return &CreateSceneCollectionParams{}
}
func (o *CreateSceneCollectionParams) WithSceneCollectionName(x string) *CreateSceneCollectionParams {
	o.SceneCollectionName = &x
	return o
}

// Returns the associated request.
func (o *CreateSceneCollectionParams) GetRequestName() string {
	return "CreateSceneCollection"
}

// Represents the response body for the CreateSceneCollection request.
type CreateSceneCollectionResponse struct {
	_response
}

/*
Creates a new scene collection, switching to it in the process.

Note: This will block until the collection has finished changing.
*/
func (c *Client) CreateSceneCollection(params *CreateSceneCollectionParams) (*CreateSceneCollectionResponse, error) {
	data := &CreateSceneCollectionResponse{}
	return data, c.client.SendRequest(params, data)
}
