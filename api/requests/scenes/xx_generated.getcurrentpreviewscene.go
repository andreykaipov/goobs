// This file has been automatically generated. Don't edit it.

package scenes

/*
GetCurrentPreviewSceneParams represents the params body for the "GetCurrentPreviewScene" request.
Gets the current preview scene.

Only available when studio mode is enabled.
*/
type GetCurrentPreviewSceneParams struct{}

/*
GetCurrentPreviewSceneResponse represents the response body for the "GetCurrentPreviewScene" request.
Gets the current preview scene.

Only available when studio mode is enabled.
*/
type GetCurrentPreviewSceneResponse struct {
	// Current preview scene
	CurrentPreviewSceneName string `json:"currentPreviewSceneName,omitempty"`
}

// GetCurrentPreviewScene sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentPreviewScene(
	paramss ...*GetCurrentPreviewSceneParams,
) (*GetCurrentPreviewSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentPreviewSceneParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("GetCurrentPreviewScene", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetCurrentPreviewSceneResponse), nil
}
