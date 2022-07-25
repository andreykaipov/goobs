// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentSceneTransitionSettingsParams represents the params body for the "SetCurrentSceneTransitionSettings" request.
Sets the settings of the current scene transition.
*/
type SetCurrentSceneTransitionSettingsParams struct {
	requests.ParamsBasic

	// Whether to overlay over the current settings or replace them
	Overlay bool `json:"overlay,omitempty"`

	// Settings object to apply to the transition. Can be `{}`
	TransitionSettings interface{} `json:"transitionSettings,omitempty"`
}

// GetSelfName just returns "SetCurrentSceneTransitionSettings".
func (o *SetCurrentSceneTransitionSettingsParams) GetSelfName() string {
	return "SetCurrentSceneTransitionSettings"
}

/*
SetCurrentSceneTransitionSettingsResponse represents the response body for the "SetCurrentSceneTransitionSettings" request.
Sets the settings of the current scene transition.
*/
type SetCurrentSceneTransitionSettingsResponse struct {
	requests.ResponseBasic
}

// SetCurrentSceneTransitionSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentSceneTransitionSettings(
	params *SetCurrentSceneTransitionSettingsParams,
) (*SetCurrentSceneTransitionSettingsResponse, error) {
	data := &SetCurrentSceneTransitionSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
