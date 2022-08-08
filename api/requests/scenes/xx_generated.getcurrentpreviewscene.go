// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the GetCurrentPreviewScene request.
type GetCurrentPreviewSceneParams struct{}

// Returns the associated request.
func (o *GetCurrentPreviewSceneParams) GetRequestName() string {
	return "GetCurrentPreviewScene"
}

// Represents the response body for the GetCurrentPreviewScene request.
type GetCurrentPreviewSceneResponse struct {
	// Current preview scene
	CurrentPreviewSceneName string `json:"currentPreviewSceneName,omitempty"`
}

/*
Gets the current preview scene.

Only available when studio mode is enabled.
*/
func (c *Client) GetCurrentPreviewScene(
	paramss ...*GetCurrentPreviewSceneParams,
) (*GetCurrentPreviewSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentPreviewSceneParams{{}}
	}
	params := paramss[0]
	data := &GetCurrentPreviewSceneResponse{}
	return data, c.SendRequest(params, data)
}
