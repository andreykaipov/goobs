// This file has been automatically generated. Don't edit it.

package config

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentProfileParams represents the params body for the "SetCurrentProfile" request.
Switches to a profile.
*/
type SetCurrentProfileParams struct {
	requests.ParamsBasic

	// Name of the profile to switch to
	ProfileName string `json:"profileName,omitempty"`
}

// GetSelfName just returns "SetCurrentProfile".
func (o *SetCurrentProfileParams) GetSelfName() string {
	return "SetCurrentProfile"
}

/*
SetCurrentProfileResponse represents the response body for the "SetCurrentProfile" request.
Switches to a profile.
*/
type SetCurrentProfileResponse struct {
	requests.ResponseBasic
}

// SetCurrentProfile sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentProfile(params *SetCurrentProfileParams) (*SetCurrentProfileResponse, error) {
	data := &SetCurrentProfileResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
