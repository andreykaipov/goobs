// This file has been automatically generated. Don't edit it.

package config

import requests "github.com/andreykaipov/goobs/api/requests"

/*
RemoveProfileParams represents the params body for the "RemoveProfile" request.
Removes a profile. If the current profile is chosen, it will change to a different profile first.
*/
type RemoveProfileParams struct {
	requests.ParamsBasic

	// Name of the profile to remove
	ProfileName string `json:"profileName,omitempty"`
}

// GetSelfName just returns "RemoveProfile".
func (o *RemoveProfileParams) GetSelfName() string {
	return "RemoveProfile"
}

/*
RemoveProfileResponse represents the response body for the "RemoveProfile" request.
Removes a profile. If the current profile is chosen, it will change to a different profile first.
*/
type RemoveProfileResponse struct {
	requests.ResponseBasic
}

// RemoveProfile sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RemoveProfile(params *RemoveProfileParams) (*RemoveProfileResponse, error) {
	data := &RemoveProfileResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
