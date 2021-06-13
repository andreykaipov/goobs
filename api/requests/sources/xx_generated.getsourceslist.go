// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourcesListParams represents the params body for the "GetSourcesList" request.
List all sources available in the running OBS instance

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#GetSourcesList.
*/
type GetSourcesListParams struct {
	requests.ParamsBasic
}

// Name just returns "GetSourcesList".
func (o *GetSourcesListParams) Name() string {
	return "GetSourcesList"
}

/*
GetSourcesListResponse represents the response body for the "GetSourcesList" request.
List all sources available in the running OBS instance

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#GetSourcesList.
*/
type GetSourcesListResponse struct {
	requests.ResponseBasic

	Sources []struct {
		// Unique source name
		Name string `json:"name"`

		// Source type. Value is one of the following: "input", "filter", "transition", "scene" or
		// "unknown"
		Type string `json:"type"`

		// Non-unique source internal type (a.k.a type id)
		TypeId string `json:"typeId"`
	} `json:"sources"`
}

// GetSourcesList sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
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
