// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the GetSceneSceneTransitionOverride request.
type GetSceneSceneTransitionOverrideParams struct {
	// Name of the scene
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *GetSceneSceneTransitionOverrideParams) GetRequestName() string {
	return "GetSceneSceneTransitionOverride"
}

// Represents the response body for the GetSceneSceneTransitionOverride request.
type GetSceneSceneTransitionOverrideResponse struct {
	// Duration of the overridden scene transition, else `null`
	TransitionDuration float64 `json:"transitionDuration,omitempty"`

	// Name of the overridden scene transition, else `null`
	TransitionName string `json:"transitionName,omitempty"`
}

// Gets the scene transition overridden for a scene.
func (c *Client) GetSceneSceneTransitionOverride(
	params *GetSceneSceneTransitionOverrideParams,
) (*GetSceneSceneTransitionOverrideResponse, error) {
	data := &GetSceneSceneTransitionOverrideResponse{}
	return data, c.SendRequest(params, data)
}
