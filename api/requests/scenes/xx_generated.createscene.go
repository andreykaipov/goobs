// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the CreateScene request.
type CreateSceneParams struct {
	// Name for the new scene
	SceneName *string `json:"sceneName,omitempty"`
}

func NewCreateSceneParams() *CreateSceneParams {
	return &CreateSceneParams{}
}
func (o *CreateSceneParams) WithSceneName(x string) *CreateSceneParams {
	o.SceneName = &x
	return o
}

// Returns the associated request.
func (o *CreateSceneParams) GetRequestName() string {
	return "CreateScene"
}

// Represents the response body for the CreateScene request.
type CreateSceneResponse struct {
	_response

	// UUID of the created scene
	SceneUuid string `json:"sceneUuid,omitempty"`
}

// Creates a new scene in OBS.
func (c *Client) CreateScene(params *CreateSceneParams) (*CreateSceneResponse, error) {
	data := &CreateSceneResponse{}
	return data, c.client.SendRequest(params, data)
}
