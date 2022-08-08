// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the GetProfileList request.
type GetProfileListParams struct{}

// Returns the associated request.
func (o *GetProfileListParams) GetRequestName() string {
	return "GetProfileList"
}

// Represents the response body for the GetProfileList request.
type GetProfileListResponse struct {
	// The name of the current profile
	CurrentProfileName string `json:"currentProfileName,omitempty"`

	// Array of all available profiles
	Profiles []string `json:"profiles,omitempty"`
}

// Gets an array of all profiles
func (c *Client) GetProfileList(paramss ...*GetProfileListParams) (*GetProfileListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetProfileListParams{{}}
	}
	params := paramss[0]
	data := &GetProfileListResponse{}
	return data, c.SendRequest(params, data)
}
