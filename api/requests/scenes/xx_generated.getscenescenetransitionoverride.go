// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the GetSceneSceneTransitionOverride request.
type GetSceneSceneTransitionOverrideParams struct {
	// Name of the scene
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewGetSceneSceneTransitionOverrideParams() *GetSceneSceneTransitionOverrideParams {
	return &GetSceneSceneTransitionOverrideParams{}
}
func (o *GetSceneSceneTransitionOverrideParams) WithSceneName(x string) *GetSceneSceneTransitionOverrideParams {
	o.SceneName = &x
	return o
}
func (o *GetSceneSceneTransitionOverrideParams) WithSceneUuid(x string) *GetSceneSceneTransitionOverrideParams {
	o.SceneUuid = &x
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

/*
Gets the scene transition overridden for a scene.

Note: A transition UUID response field is not currently able to be implemented as of 2024-1-18.
*/
func (c *Client) GetSceneSceneTransitionOverride(
	paramss ...*GetSceneSceneTransitionOverrideParams,
) (*GetSceneSceneTransitionOverrideResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSceneSceneTransitionOverrideParams{{}}
	}
	params := paramss[0]
	data := &GetSceneSceneTransitionOverrideResponse{}
	return data, c.client.SendRequest(params, data)
}
