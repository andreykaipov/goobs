// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the SetSceneName request.
type SetSceneNameParams struct {
	// New name for the scene
	NewSceneName string `json:"newSceneName,omitempty"`

	// Name of the scene to be renamed
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *SetSceneNameParams) GetRequestName() string {
	return "SetSceneName"
}

// Represents the response body for the SetSceneName request.
type SetSceneNameResponse struct{}

// Sets the name of a scene (rename).
func (c *Client) SetSceneName(params *SetSceneNameParams) (*SetSceneNameResponse, error) {
	data := &SetSceneNameResponse{}
	return data, c.SendRequest(params, data)
}
