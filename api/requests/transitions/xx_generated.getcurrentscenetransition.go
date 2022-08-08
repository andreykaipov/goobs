// This file has been automatically generated. Don't edit it.

package transitions

// Represents the request body for the GetCurrentSceneTransition request.
type GetCurrentSceneTransitionParams struct{}

// Returns the associated request.
func (o *GetCurrentSceneTransitionParams) GetRequestName() string {
	return "GetCurrentSceneTransition"
}

// Represents the response body for the GetCurrentSceneTransition request.
type GetCurrentSceneTransitionResponse struct {
	// Whether the transition supports being configured
	TransitionConfigurable bool `json:"transitionConfigurable,omitempty"`

	// Configured transition duration in milliseconds. `null` if transition is fixed
	TransitionDuration float64 `json:"transitionDuration,omitempty"`

	// Whether the transition uses a fixed (unconfigurable) duration
	TransitionFixed bool `json:"transitionFixed,omitempty"`

	// Kind of the transition
	TransitionKind string `json:"transitionKind,omitempty"`

	// Name of the transition
	TransitionName string `json:"transitionName,omitempty"`

	// Object of settings for the transition. `null` if transition is not configurable
	TransitionSettings map[string]interface{} `json:"transitionSettings,omitempty"`
}

// Gets information about the current scene transition.
func (c *Client) GetCurrentSceneTransition(
	paramss ...*GetCurrentSceneTransitionParams,
) (*GetCurrentSceneTransitionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentSceneTransitionParams{{}}
	}
	params := paramss[0]
	data := &GetCurrentSceneTransitionResponse{}
	return data, c.SendRequest(params, data)
}
