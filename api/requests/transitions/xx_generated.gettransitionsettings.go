// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetTransitionSettingsParams represents the params body for the "GetTransitionSettings" request.
Get the current settings of a transition
Since 4.9.0.
*/
type GetTransitionSettingsParams struct {
	requests.ParamsBasic

	// Transition name
	TransitionName string `json:"transitionName,omitempty"`
}

// GetSelfName just returns "GetTransitionSettings".
func (o *GetTransitionSettingsParams) GetSelfName() string {
	return "GetTransitionSettings"
}

/*
GetTransitionSettingsResponse represents the response body for the "GetTransitionSettings" request.
Get the current settings of a transition
Since v4.9.0.
*/
type GetTransitionSettingsResponse struct {
	requests.ResponseBasic

	// Current transition settings
	TransitionSettings map[string]interface{} `json:"transitionSettings,omitempty"`
}

// GetTransitionSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetTransitionSettings(params *GetTransitionSettingsParams) (*GetTransitionSettingsResponse, error) {
	data := &GetTransitionSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
