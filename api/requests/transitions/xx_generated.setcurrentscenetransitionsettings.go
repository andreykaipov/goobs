// This file has been automatically generated. Don't edit it.

package transitions

// Represents the request body for the SetCurrentSceneTransitionSettings request.
type SetCurrentSceneTransitionSettingsParams struct {
	// Whether to overlay over the current settings or replace them
	Overlay *bool `json:"overlay,omitempty"`

	// Settings object to apply to the transition. Can be `{}`
	TransitionSettings map[string]interface{} `json:"transitionSettings,omitempty"`
}

// Returns the associated request.
func (o *SetCurrentSceneTransitionSettingsParams) GetRequestName() string {
	return "SetCurrentSceneTransitionSettings"
}

// Represents the response body for the SetCurrentSceneTransitionSettings request.
type SetCurrentSceneTransitionSettingsResponse struct{}

// Sets the settings of the current scene transition.
func (c *Client) SetCurrentSceneTransitionSettings(
	params *SetCurrentSceneTransitionSettingsParams,
) (*SetCurrentSceneTransitionSettingsResponse, error) {
	data := &SetCurrentSceneTransitionSettingsResponse{}
	return data, c.SendRequest(params, data)
}
