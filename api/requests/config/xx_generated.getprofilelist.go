// This file has been automatically generated. Don't edit it.

package config

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetProfileListParams represents the params body for the "GetProfileList" request.
Gets an array of all profiles
*/
type GetProfileListParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetProfileList".
func (o *GetProfileListParams) GetSelfName() string {
	return "GetProfileList"
}

/*
GetProfileListResponse represents the response body for the "GetProfileList" request.
Gets an array of all profiles
*/
type GetProfileListResponse struct {
	requests.ResponseBasic

	// The name of the current profile
	CurrentProfileName string `json:"currentProfileName,omitempty"`

	// Array of all available profiles
	Profiles []string `json:"profiles,omitempty"`
}

// GetProfileList sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) GetProfileList(paramss ...*GetProfileListParams) (*GetProfileListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetProfileListParams{{}}
	}
	params := paramss[0]
	data := &GetProfileListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
