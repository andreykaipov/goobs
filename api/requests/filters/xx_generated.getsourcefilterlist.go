// This file has been automatically generated. Don't edit it.

package filters

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourceFilterListParams represents the params body for the "GetSourceFilterList" request.
Gets an array of all of a source's filters.
*/
type GetSourceFilterListParams struct {
	requests.ParamsBasic

	// Name of the source
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "GetSourceFilterList".
func (o *GetSourceFilterListParams) GetSelfName() string {
	return "GetSourceFilterList"
}

/*
GetSourceFilterListResponse represents the response body for the "GetSourceFilterList" request.
Gets an array of all of a source's filters.
*/
type GetSourceFilterListResponse struct {
	requests.ResponseBasic

	// Array of filters
	Filters []interface{} `json:"filters,omitempty"`
}

// GetSourceFilterList sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceFilterList(params *GetSourceFilterListParams) (*GetSourceFilterListResponse, error) {
	data := &GetSourceFilterListResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
