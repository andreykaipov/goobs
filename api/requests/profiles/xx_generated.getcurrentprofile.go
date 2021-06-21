// This file has been automatically generated. Don't edit it.

package profiles

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetCurrentProfileParams represents the params body for the "GetCurrentProfile" request.
Get the name of the current profile.
Since 4.0.0.
*/
type GetCurrentProfileParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetCurrentProfile".
func (o *GetCurrentProfileParams) GetSelfName() string {
	return "GetCurrentProfile"
}

/*
GetCurrentProfileResponse represents the response body for the "GetCurrentProfile" request.
Get the name of the current profile.
Since v4.0.0.
*/
type GetCurrentProfileResponse struct {
	requests.ResponseBasic

	// Name of the currently active profile.
	ProfileName string `json:"profile-name,omitempty"`
}

// GetCurrentProfile sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) GetCurrentProfile(paramss ...*GetCurrentProfileParams) (*GetCurrentProfileResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentProfileParams{{}}
	}
	params := paramss[0]
	data := &GetCurrentProfileResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
