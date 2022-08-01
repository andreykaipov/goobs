// This file has been automatically generated. Don't edit it.

package scenes

/*
RemoveSceneParams represents the params body for the "RemoveScene" request.
Removes a scene from OBS.
*/
type RemoveSceneParams struct {
	// Name of the scene to remove
	SceneName string `json:"sceneName,omitempty"`
}

/*
RemoveSceneResponse represents the response body for the "RemoveScene" request.
Removes a scene from OBS.
*/
type RemoveSceneResponse struct{}

// RemoveScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RemoveScene(params *RemoveSceneParams) (*RemoveSceneResponse, error) {
	resp, err := c.SendRequest("RemoveScene", params)
	if err != nil {
		return nil, err
	}
	return resp.(*RemoveSceneResponse), nil
}
