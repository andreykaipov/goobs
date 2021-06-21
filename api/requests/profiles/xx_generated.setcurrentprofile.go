// This file has been automatically generated. Don't edit it.

package profiles

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentProfileParams represents the params body for the "SetCurrentProfile" request.
Set the currently active profile.
Since 4.0.0.
*/
type SetCurrentProfileParams struct {
	requests.ParamsBasic

	// Name of the desired profile.
	ProfileName string `json:"profile-name,omitempty"`
}

// GetSelfName just returns "SetCurrentProfile".
func (o *SetCurrentProfileParams) GetSelfName() string {
	return "SetCurrentProfile"
}

/*
SetCurrentProfileResponse represents the response body for the "SetCurrentProfile" request.
Set the currently active profile.
Since v4.0.0.
*/
type SetCurrentProfileResponse struct {
	requests.ResponseBasic
}

// SetCurrentProfile sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentProfile(params *SetCurrentProfileParams) (*SetCurrentProfileResponse, error) {
	data := &SetCurrentProfileResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
