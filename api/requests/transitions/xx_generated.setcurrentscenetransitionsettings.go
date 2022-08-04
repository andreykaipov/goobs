// This file has been automatically generated. Don't edit it.

package transitions

/*
SetCurrentSceneTransitionSettingsParams represents the params body for the "SetCurrentSceneTransitionSettings" request.
Sets the settings of the current scene transition.
*/
type SetCurrentSceneTransitionSettingsParams struct {
	// Whether to overlay over the current settings or replace them
	Overlay *bool `json:"overlay,omitempty"`

	// Settings object to apply to the transition. Can be `{}`
	TransitionSettings interface{} `json:"transitionSettings,omitempty"`
}

/*
SetCurrentSceneTransitionSettingsResponse represents the response body for the "SetCurrentSceneTransitionSettings" request.
Sets the settings of the current scene transition.
*/
type SetCurrentSceneTransitionSettingsResponse struct{}

// SetCurrentSceneTransitionSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentSceneTransitionSettings(
	params *SetCurrentSceneTransitionSettingsParams,
) (*SetCurrentSceneTransitionSettingsResponse, error) {
	resp, err := c.SendRequest("SetCurrentSceneTransitionSettings", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetCurrentSceneTransitionSettingsResponse), nil
}
