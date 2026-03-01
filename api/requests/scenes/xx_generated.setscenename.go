// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the SetSceneName request.
type SetSceneNameParams struct {
	// UUID of the canvas the scene is in, if using the sceneName field
	CanvasUuid *string `json:"canvasUuid,omitempty"`

	// New name for the scene
	NewSceneName *string `json:"newSceneName,omitempty"`

	// Name of the scene to be renamed
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene to be renamed
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewSetSceneNameParams() *SetSceneNameParams {
	return &SetSceneNameParams{}
}
func (o *SetSceneNameParams) WithCanvasUuid(x string) *SetSceneNameParams {
	o.CanvasUuid = &x
	return o
}
func (o *SetSceneNameParams) WithNewSceneName(x string) *SetSceneNameParams {
	o.NewSceneName = &x
	return o
}
func (o *SetSceneNameParams) WithSceneName(x string) *SetSceneNameParams {
	o.SceneName = &x
	return o
}
func (o *SetSceneNameParams) WithSceneUuid(x string) *SetSceneNameParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *SetSceneNameParams) GetRequestName() string {
	return "SetSceneName"
}

// Represents the response body for the SetSceneName request.
type SetSceneNameResponse struct {
	_response
}

// Sets the name of a scene (rename).
func (c *Client) SetSceneName(params *SetSceneNameParams) (*SetSceneNameResponse, error) {
	data := &SetSceneNameResponse{}
	return data, c.client.SendRequest(params, data)
}
