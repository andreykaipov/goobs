// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the GetSceneSceneTransitionOverride request.
type GetSceneSceneTransitionOverrideParams struct {
	// Name of the scene
	SceneName *string `json:"sceneName,omitempty"`
}

func NewGetSceneSceneTransitionOverrideParams() *GetSceneSceneTransitionOverrideParams {
	return &GetSceneSceneTransitionOverrideParams{}
}
func (o *GetSceneSceneTransitionOverrideParams) WithSceneName(x string) *GetSceneSceneTransitionOverrideParams {
	o.SceneName = &x
	return o
}

// Returns the associated request.
func (o *GetSceneSceneTransitionOverrideParams) GetRequestName() string {
	return "GetSceneSceneTransitionOverride"
}

// Represents the response body for the GetSceneSceneTransitionOverride request.
type GetSceneSceneTransitionOverrideResponse struct {
	_response

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
	return data, c.client.SendRequest(params, data)
}
