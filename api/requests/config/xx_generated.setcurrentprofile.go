// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the SetCurrentProfile request.
type SetCurrentProfileParams struct {
	// Name of the profile to switch to
	ProfileName string `json:"profileName,omitempty"`
}

// Returns the associated request.
func (o *SetCurrentProfileParams) GetRequestName() string {
	return "SetCurrentProfile"
}

// Represents the response body for the SetCurrentProfile request.
type SetCurrentProfileResponse struct{}

// Switches to a profile.
func (c *Client) SetCurrentProfile(params *SetCurrentProfileParams) (*SetCurrentProfileResponse, error) {
	data := &SetCurrentProfileResponse{}
	return data, c.SendRequest(params, data)
}
