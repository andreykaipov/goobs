// This file has been automatically generated. Don't edit it.

package profiles

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ListProfilesParams represents the params body for the "ListProfiles" request.
Get a list of available profiles.
Since 4.0.0.
*/
type ListProfilesParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "ListProfiles".
func (o *ListProfilesParams) GetSelfName() string {
	return "ListProfiles"
}

/*
ListProfilesResponse represents the response body for the "ListProfiles" request.
Get a list of available profiles.
Since v4.0.0.
*/
type ListProfilesResponse struct {
	requests.ResponseBasic

	Profiles []*Profile `json:"profiles,omitempty"`
}

type Profile struct {
	// Filter name
	ProfileName string `json:"profile-name,omitempty"`
}

// ListProfiles sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) ListProfiles(paramss ...*ListProfilesParams) (*ListProfilesResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ListProfilesParams{{}}
	}
	params := paramss[0]
	data := &ListProfilesResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
