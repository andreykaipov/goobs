// This file has been automatically generated. Don't edit it.

package config

/*
SetCurrentProfileParams represents the params body for the "SetCurrentProfile" request.
Switches to a profile.
*/
type SetCurrentProfileParams struct {
	// Name of the profile to switch to
	ProfileName string `json:"profileName,omitempty"`
}

/*
SetCurrentProfileResponse represents the response body for the "SetCurrentProfile" request.
Switches to a profile.
*/
type SetCurrentProfileResponse struct{}

// SetCurrentProfile sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentProfile(params *SetCurrentProfileParams) (*SetCurrentProfileResponse, error) {
	resp, err := c.SendRequest("SetCurrentProfile", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetCurrentProfileResponse), nil
}
