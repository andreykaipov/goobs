// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the SetSceneSceneTransitionOverride request.
type SetSceneSceneTransitionOverrideParams struct {
	// Name of the scene
	SceneName string `json:"sceneName,omitempty"`

	// Duration to use for any overridden transition. Specify `null` to remove
	TransitionDuration float64 `json:"transitionDuration,omitempty"`

	// Name of the scene transition to use as override. Specify `null` to remove
	TransitionName string `json:"transitionName,omitempty"`
}

// Returns the associated request.
func (o *SetSceneSceneTransitionOverrideParams) GetRequestName() string {
	return "SetSceneSceneTransitionOverride"
}

// Represents the response body for the SetSceneSceneTransitionOverride request.
type SetSceneSceneTransitionOverrideResponse struct{}

// Gets the scene transition overridden for a scene.
func (c *Client) SetSceneSceneTransitionOverride(
	params *SetSceneSceneTransitionOverrideParams,
) (*SetSceneSceneTransitionOverrideResponse, error) {
	data := &SetSceneSceneTransitionOverrideResponse{}
	return data, c.SendRequest(params, data)
}
