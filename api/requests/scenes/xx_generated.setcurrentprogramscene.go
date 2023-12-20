// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the SetCurrentProgramScene request.
type SetCurrentProgramSceneParams struct {
	// Scene to set as the current program scene
	SceneName *string `json:"sceneName,omitempty"`
}

func NewSetCurrentProgramSceneParams() *SetCurrentProgramSceneParams {
	return &SetCurrentProgramSceneParams{}
}
func (o *SetCurrentProgramSceneParams) WithSceneName(x string) *SetCurrentProgramSceneParams {
	o.SceneName = &x
	return o
}

// Returns the associated request.
func (o *SetCurrentProgramSceneParams) GetRequestName() string {
	return "SetCurrentProgramScene"
}

// Represents the response body for the SetCurrentProgramScene request.
type SetCurrentProgramSceneResponse struct {
	_response
}

// Sets the current program scene.
func (c *Client) SetCurrentProgramScene(params *SetCurrentProgramSceneParams) (*SetCurrentProgramSceneResponse, error) {
	data := &SetCurrentProgramSceneResponse{}
	return data, c.client.SendRequest(params, data)
}
