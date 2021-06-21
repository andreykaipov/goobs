// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourcesListParams represents the params body for the "GetSourcesList" request.
List all sources available in the running OBS instance
Since 4.3.0.
*/
type GetSourcesListParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetSourcesList".
func (o *GetSourcesListParams) GetSelfName() string {
	return "GetSourcesList"
}

/*
GetSourcesListResponse represents the response body for the "GetSourcesList" request.
List all sources available in the running OBS instance
Since v4.3.0.
*/
type GetSourcesListResponse struct {
	requests.ResponseBasic

	Sources []*Source `json:"sources,omitempty"`
}

type Source struct {
	// Unique source name
	Name string `json:"name,omitempty"`

	// Source type. Value is one of the following: "input", "filter", "transition", "scene" or "unknown"
	Type string `json:"type,omitempty"`

	// Non-unique source internal type (a.k.a kind)
	TypeId string `json:"typeId,omitempty"`
}

// GetSourcesList sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) GetSourcesList(paramss ...*GetSourcesListParams) (*GetSourcesListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSourcesListParams{{}}
	}
	params := paramss[0]
	data := &GetSourcesListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
