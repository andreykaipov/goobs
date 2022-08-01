// This file has been automatically generated. Don't edit it.

package scenes

/*
SetCurrentProgramSceneParams represents the params body for the "SetCurrentProgramScene" request.
Sets the current program scene.
*/
type SetCurrentProgramSceneParams struct {
	// Scene to set as the current program scene
	SceneName string `json:"sceneName,omitempty"`
}

/*
SetCurrentProgramSceneResponse represents the response body for the "SetCurrentProgramScene" request.
Sets the current program scene.
*/
type SetCurrentProgramSceneResponse struct{}

// SetCurrentProgramScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentProgramScene(params *SetCurrentProgramSceneParams) (*SetCurrentProgramSceneResponse, error) {
	resp, err := c.SendRequest("SetCurrentProgramScene", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetCurrentProgramSceneResponse), nil
}
