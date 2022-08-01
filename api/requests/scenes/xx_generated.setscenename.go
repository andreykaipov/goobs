// This file has been automatically generated. Don't edit it.

package scenes

/*
SetSceneNameParams represents the params body for the "SetSceneName" request.
Sets the name of a scene (rename).
*/
type SetSceneNameParams struct {
	// New name for the scene
	NewSceneName string `json:"newSceneName,omitempty"`

	// Name of the scene to be renamed
	SceneName string `json:"sceneName,omitempty"`
}

/*
SetSceneNameResponse represents the response body for the "SetSceneName" request.
Sets the name of a scene (rename).
*/
type SetSceneNameResponse struct{}

// SetSceneName sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneName(params *SetSceneNameParams) (*SetSceneNameResponse, error) {
	resp, err := c.SendRequest("SetSceneName", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetSceneNameResponse), nil
}
