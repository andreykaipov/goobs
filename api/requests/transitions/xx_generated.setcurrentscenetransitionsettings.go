// This file has been automatically generated. Don't edit it.

package transitions

// Represents the request body for the SetCurrentSceneTransitionSettings request.
type SetCurrentSceneTransitionSettingsParams struct {
	// Whether to overlay over the current settings or replace them
	Overlay *bool `json:"overlay,omitempty"`

	// Settings object to apply to the transition. Can be `{}`
	TransitionSettings map[string]any `json:"transitionSettings,omitempty"`
}

func NewSetCurrentSceneTransitionSettingsParams() *SetCurrentSceneTransitionSettingsParams {
	return &SetCurrentSceneTransitionSettingsParams{}
}
func (o *SetCurrentSceneTransitionSettingsParams) WithOverlay(x bool) *SetCurrentSceneTransitionSettingsParams {
	o.Overlay = &x
	return o
}

func (o *SetCurrentSceneTransitionSettingsParams) WithTransitionSettings(
	x map[string]any,
) *SetCurrentSceneTransitionSettingsParams {
	o.TransitionSettings = x
	return o
}

// Returns the associated request.
func (o *SetCurrentSceneTransitionSettingsParams) GetRequestName() string {
	return "SetCurrentSceneTransitionSettings"
}

// Represents the response body for the SetCurrentSceneTransitionSettings request.
type SetCurrentSceneTransitionSettingsResponse struct {
	_response
}

// Sets the settings of the current scene transition.
func (c *Client) SetCurrentSceneTransitionSettings(
	params *SetCurrentSceneTransitionSettingsParams,
) (*SetCurrentSceneTransitionSettingsResponse, error) {
	data := &SetCurrentSceneTransitionSettingsResponse{}
	return data, c.client.SendRequest(params, data)
}
