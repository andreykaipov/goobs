// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the SetCurrentPreviewScene request.
type SetCurrentPreviewSceneParams struct {
	// Scene name to set as the current preview scene
	SceneName *string `json:"sceneName,omitempty"`

	// Scene UUID to set as the current preview scene
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewSetCurrentPreviewSceneParams() *SetCurrentPreviewSceneParams {
	return &SetCurrentPreviewSceneParams{}
}
func (o *SetCurrentPreviewSceneParams) WithSceneName(x string) *SetCurrentPreviewSceneParams {
	o.SceneName = &x
	return o
}
func (o *SetCurrentPreviewSceneParams) WithSceneUuid(x string) *SetCurrentPreviewSceneParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *SetCurrentPreviewSceneParams) GetRequestName() string {
	return "SetCurrentPreviewScene"
}

// Represents the response body for the SetCurrentPreviewScene request.
type SetCurrentPreviewSceneResponse struct {
	_response
}

/*
Sets the current preview scene.

Only available when studio mode is enabled.
*/
func (c *Client) SetCurrentPreviewScene(
	paramss ...*SetCurrentPreviewSceneParams,
) (*SetCurrentPreviewSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*SetCurrentPreviewSceneParams{{}}
	}
	params := paramss[0]
	data := &SetCurrentPreviewSceneResponse{}
	return data, c.client.SendRequest(params, data)
}
