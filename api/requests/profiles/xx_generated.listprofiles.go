// This file has been automatically generated. Don't edit it.

package profiles

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ListProfilesParams represents the params body for the "ListProfiles" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListProfiles.
*/
type ListProfilesParams struct {
	requests.ParamsBasic
}

// Name just returns "ListProfiles".
func (o *ListProfilesParams) Name() string {
	return "ListProfiles"
}

/*
ListProfilesResponse represents the response body for the "ListProfiles" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ListProfiles.
*/
type ListProfilesResponse struct {
	requests.ResponseBasic

	// List of available profiles.
	Profiles []map[string]interface{} `json:"profiles"`
}

// ListProfiles sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) ListProfiles(paramss ...*ListProfilesParams) (*ListProfilesResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ListProfilesParams{{}}
	}
	params := paramss[0]
	data := &ListProfilesResponse{}
	if err := c.WriteMessage(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
