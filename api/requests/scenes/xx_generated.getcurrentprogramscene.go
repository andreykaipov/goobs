// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the GetCurrentProgramScene request.
type GetCurrentProgramSceneParams struct{}

// Returns the associated request.
func (o *GetCurrentProgramSceneParams) GetRequestName() string {
	return "GetCurrentProgramScene"
}

// Represents the response body for the GetCurrentProgramScene request.
type GetCurrentProgramSceneResponse struct {
	_response

	// Current program scene name (Deprecated)
	CurrentProgramSceneName string `json:"currentProgramSceneName,omitempty"`

	// Current program scene UUID (Deprecated)
	CurrentProgramSceneUuid string `json:"currentProgramSceneUuid,omitempty"`

	// Current program scene name
	SceneName string `json:"sceneName,omitempty"`

	// Current program scene UUID
	SceneUuid string `json:"sceneUuid,omitempty"`
}

/*
Gets the current program scene.

Note: This request is slated to have the `currentProgram`-prefixed fields removed from in an upcoming RPC version.
*/
func (c *Client) GetCurrentProgramScene(
	paramss ...*GetCurrentProgramSceneParams,
) (*GetCurrentProgramSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentProgramSceneParams{{}}
	}
	params := paramss[0]
	data := &GetCurrentProgramSceneResponse{}
	return data, c.client.SendRequest(params, data)
}
