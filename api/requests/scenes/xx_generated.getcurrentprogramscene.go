// This file has been automatically generated. Don't edit it.

package scenes

/*
GetCurrentProgramSceneParams represents the params body for the "GetCurrentProgramScene" request.
Gets the current program scene.
*/
type GetCurrentProgramSceneParams struct{}

/*
GetCurrentProgramSceneResponse represents the response body for the "GetCurrentProgramScene" request.
Gets the current program scene.
*/
type GetCurrentProgramSceneResponse struct {
	// Current program scene
	CurrentProgramSceneName string `json:"currentProgramSceneName,omitempty"`
}

// GetCurrentProgramScene sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentProgramScene(
	paramss ...*GetCurrentProgramSceneParams,
) (*GetCurrentProgramSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentProgramSceneParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("GetCurrentProgramScene", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetCurrentProgramSceneResponse), nil
}
