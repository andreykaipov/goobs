// This file has been automatically generated. Don't edit it.

package config

import requests "github.com/andreykaipov/goobs/api/requests"

/*
CreateProfileParams represents the params body for the "CreateProfile" request.
Creates a new profile, switching to it in the process
*/
type CreateProfileParams struct {
	requests.ParamsBasic

	// Name for the new profile
	ProfileName string `json:"profileName,omitempty"`
}

// GetSelfName just returns "CreateProfile".
func (o *CreateProfileParams) GetSelfName() string {
	return "CreateProfile"
}

/*
CreateProfileResponse represents the response body for the "CreateProfile" request.
Creates a new profile, switching to it in the process
*/
type CreateProfileResponse struct {
	requests.ResponseBasic
}

// CreateProfile sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) CreateProfile(params *CreateProfileParams) (*CreateProfileResponse, error) {
	data := &CreateProfileResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
