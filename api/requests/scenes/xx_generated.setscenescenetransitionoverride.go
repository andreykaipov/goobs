// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the SetSceneSceneTransitionOverride request.
type SetSceneSceneTransitionOverrideParams struct {
	// Name of the scene
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene
	SceneUuid *string `json:"sceneUuid,omitempty"`

	// Duration to use for any overridden transition. Specify `null` to remove
	TransitionDuration *float64 `json:"transitionDuration,omitempty"`

	// Name of the scene transition to use as override. Specify `null` to remove
	TransitionName *string `json:"transitionName,omitempty"`
}

func NewSetSceneSceneTransitionOverrideParams() *SetSceneSceneTransitionOverrideParams {
	return &SetSceneSceneTransitionOverrideParams{}
}
func (o *SetSceneSceneTransitionOverrideParams) WithSceneName(x string) *SetSceneSceneTransitionOverrideParams {
	o.SceneName = &x
	return o
}
func (o *SetSceneSceneTransitionOverrideParams) WithSceneUuid(x string) *SetSceneSceneTransitionOverrideParams {
	o.SceneUuid = &x
	return o
}

func (o *SetSceneSceneTransitionOverrideParams) WithTransitionDuration(
	x float64,
) *SetSceneSceneTransitionOverrideParams {
	o.TransitionDuration = &x
	return o
}
func (o *SetSceneSceneTransitionOverrideParams) WithTransitionName(x string) *SetSceneSceneTransitionOverrideParams {
	o.TransitionName = &x
	return o
}

// Returns the associated request.
func (o *SetSceneSceneTransitionOverrideParams) GetRequestName() string {
	return "SetSceneSceneTransitionOverride"
}

// Represents the response body for the SetSceneSceneTransitionOverride request.
type SetSceneSceneTransitionOverrideResponse struct {
	_response
}

// Sets the scene transition overridden for a scene.
func (c *Client) SetSceneSceneTransitionOverride(
	paramss ...*SetSceneSceneTransitionOverrideParams,
) (*SetSceneSceneTransitionOverrideResponse, error) {
	if len(paramss) == 0 {
		paramss = []*SetSceneSceneTransitionOverrideParams{{}}
	}
	params := paramss[0]
	data := &SetSceneSceneTransitionOverrideResponse{}
	return data, c.client.SendRequest(params, data)
}
