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
	_response

	// Current preview scene name
	CurrentPreviewSceneName string `json:"currentPreviewSceneName,omitempty"`

	// Current preview scene UUID
	CurrentPreviewSceneUuid string `json:"currentPreviewSceneUuid,omitempty"`

	// Current preview scene name
	SceneName string `json:"sceneName,omitempty"`

	// Current preview scene UUID
	SceneUuid string `json:"sceneUuid,omitempty"`
}

/*
Gets the current preview scene.

Only available when studio mode is enabled.

Note: This request is slated to have the `currentPreview`-prefixed fields removed from in an upcoming RPC version.
*/
func (c *Client) GetCurrentPreviewScene(
	paramss ...*GetCurrentPreviewSceneParams,
) (*GetCurrentPreviewSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentPreviewSceneParams{{}}
	}
	params := paramss[0]
	data := &GetCurrentPreviewSceneResponse{}
	return data, c.client.SendRequest(params, data)
}
