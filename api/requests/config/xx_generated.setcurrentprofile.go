// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the SetCurrentProfile request.
type SetCurrentProfileParams struct {
	// Name of the profile to switch to
	ProfileName *string `json:"profileName,omitempty"`
}

func NewSetCurrentProfileParams() *SetCurrentProfileParams {
	return &SetCurrentProfileParams{}
}
func (o *SetCurrentProfileParams) WithProfileName(x string) *SetCurrentProfileParams {
	o.ProfileName = &x
	return o
}

// Returns the associated request.
func (o *SetCurrentProfileParams) GetRequestName() string {
	return "SetCurrentProfile"
}

// Represents the response body for the SetCurrentProfile request.
type SetCurrentProfileResponse struct {
	_response
}

// Switches to a profile.
func (c *Client) SetCurrentProfile(params *SetCurrentProfileParams) (*SetCurrentProfileResponse, error) {
	data := &SetCurrentProfileResponse{}
	return data, c.client.SendRequest(params, data)
}
