// This file has been automatically generated. Don't edit it.

package config

/*
RemoveProfileParams represents the params body for the "RemoveProfile" request.
Removes a profile. If the current profile is chosen, it will change to a different profile first.
*/
type RemoveProfileParams struct {
	// Name of the profile to remove
	ProfileName string `json:"profileName,omitempty"`
}

/*
RemoveProfileResponse represents the response body for the "RemoveProfile" request.
Removes a profile. If the current profile is chosen, it will change to a different profile first.
*/
type RemoveProfileResponse struct{}

// RemoveProfile sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RemoveProfile(params *RemoveProfileParams) (*RemoveProfileResponse, error) {
	resp, err := c.SendRequest("RemoveProfile", params)
	if err != nil {
		return nil, err
	}
	return resp.(*RemoveProfileResponse), nil
}
