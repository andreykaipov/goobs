// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the RemoveProfile request.
type RemoveProfileParams struct {
	// Name of the profile to remove
	ProfileName string `json:"profileName,omitempty"`
}

// Returns the associated request.
func (o *RemoveProfileParams) GetRequestName() string {
	return "RemoveProfile"
}

// Represents the response body for the RemoveProfile request.
type RemoveProfileResponse struct{}

// Removes a profile. If the current profile is chosen, it will change to a different profile first.
func (c *Client) RemoveProfile(params *RemoveProfileParams) (*RemoveProfileResponse, error) {
	data := &RemoveProfileResponse{}
	return data, c.SendRequest(params, data)
}
