// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetTransitionSettingsParams represents the params body for the "SetTransitionSettings" request.
Change the current settings of a transition
Since 4.9.0.
*/
type SetTransitionSettingsParams struct {
	requests.ParamsBasic

	// Transition name
	TransitionName string `json:"transitionName,omitempty"`

	// Transition settings (they can be partial)
	TransitionSettings map[string]interface{} `json:"transitionSettings,omitempty"`
}

// GetSelfName just returns "SetTransitionSettings".
func (o *SetTransitionSettingsParams) GetSelfName() string {
	return "SetTransitionSettings"
}

/*
SetTransitionSettingsResponse represents the response body for the "SetTransitionSettings" request.
Change the current settings of a transition
Since v4.9.0.
*/
type SetTransitionSettingsResponse struct {
	requests.ResponseBasic

	// Updated transition settings
	TransitionSettings map[string]interface{} `json:"transitionSettings,omitempty"`
}

// SetTransitionSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetTransitionSettings(params *SetTransitionSettingsParams) (*SetTransitionSettingsResponse, error) {
	data := &SetTransitionSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
