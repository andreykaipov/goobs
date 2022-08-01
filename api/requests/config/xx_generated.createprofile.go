// This file has been automatically generated. Don't edit it.

package config

/*
CreateProfileParams represents the params body for the "CreateProfile" request.
Creates a new profile, switching to it in the process
*/
type CreateProfileParams struct {
	// Name for the new profile
	ProfileName string `json:"profileName,omitempty"`
}

/*
CreateProfileResponse represents the response body for the "CreateProfile" request.
Creates a new profile, switching to it in the process
*/
type CreateProfileResponse struct{}

// CreateProfile sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) CreateProfile(params *CreateProfileParams) (*CreateProfileResponse, error) {
	resp, err := c.SendRequest("CreateProfile", params)
	if err != nil {
		return nil, err
	}
	return resp.(*CreateProfileResponse), nil
}
