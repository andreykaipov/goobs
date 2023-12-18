// This file has been automatically generated. Don't edit it.

package config

import api "github.com/andreykaipov/goobs/api"

// Represents the request body for the RemoveProfile request.
type RemoveProfileParams struct {
	// Name of the profile to remove
	ProfileName *string `json:"profileName,omitempty"`
}

func NewRemoveProfileParams() *RemoveProfileParams {
	return &RemoveProfileParams{}
}
func (o *RemoveProfileParams) WithProfileName(x string) *RemoveProfileParams {
	o.ProfileName = &x
	return o
}

// Returns the associated request.
func (o *RemoveProfileParams) GetRequestName() string {
	return "RemoveProfile"
}

// Represents the response body for the RemoveProfile request.
type RemoveProfileResponse struct {
	api.ResponseCommon
}

// Removes a profile. If the current profile is chosen, it will change to a different profile first.
func (c *Client) RemoveProfile(params *RemoveProfileParams) (*RemoveProfileResponse, error) {
	data := &RemoveProfileResponse{}
	return data, c.SendRequest(params, data)
}
