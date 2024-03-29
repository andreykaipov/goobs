// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the CreateProfile request.
type CreateProfileParams struct {
	// Name for the new profile
	ProfileName *string `json:"profileName,omitempty"`
}

func NewCreateProfileParams() *CreateProfileParams {
	return &CreateProfileParams{}
}
func (o *CreateProfileParams) WithProfileName(x string) *CreateProfileParams {
	o.ProfileName = &x
	return o
}

// Returns the associated request.
func (o *CreateProfileParams) GetRequestName() string {
	return "CreateProfile"
}

// Represents the response body for the CreateProfile request.
type CreateProfileResponse struct {
	_response
}

// Creates a new profile, switching to it in the process
func (c *Client) CreateProfile(params *CreateProfileParams) (*CreateProfileResponse, error) {
	data := &CreateProfileResponse{}
	return data, c.client.SendRequest(params, data)
}
