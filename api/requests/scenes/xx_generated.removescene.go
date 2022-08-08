// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the RemoveScene request.
type RemoveSceneParams struct {
	// Name of the scene to remove
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *RemoveSceneParams) GetRequestName() string {
	return "RemoveScene"
}

// Represents the response body for the RemoveScene request.
type RemoveSceneResponse struct{}

// Removes a scene from OBS.
func (c *Client) RemoveScene(params *RemoveSceneParams) (*RemoveSceneResponse, error) {
	data := &RemoveSceneResponse{}
	return data, c.SendRequest(params, data)
}
