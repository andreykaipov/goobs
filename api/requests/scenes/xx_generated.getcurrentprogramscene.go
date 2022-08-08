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
	// Current program scene
	CurrentProgramSceneName string `json:"currentProgramSceneName,omitempty"`
}

// Gets the current program scene.
func (c *Client) GetCurrentProgramScene(
	paramss ...*GetCurrentProgramSceneParams,
) (*GetCurrentProgramSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentProgramSceneParams{{}}
	}
	params := paramss[0]
	data := &GetCurrentProgramSceneResponse{}
	return data, c.SendRequest(params, data)
}
