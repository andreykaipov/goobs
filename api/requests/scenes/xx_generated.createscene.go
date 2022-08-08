// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the CreateScene request.
type CreateSceneParams struct {
	// Name for the new scene
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *CreateSceneParams) GetRequestName() string {
	return "CreateScene"
}

// Represents the response body for the CreateScene request.
type CreateSceneResponse struct{}

// Creates a new scene in OBS.
func (c *Client) CreateScene(params *CreateSceneParams) (*CreateSceneResponse, error) {
	data := &CreateSceneResponse{}
	return data, c.SendRequest(params, data)
}
