// This file has been automatically generated. Don't edit it.

package filters

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the GetSourceFilterList request.
type GetSourceFilterListParams struct {
	// Name of the source
	SourceName string `json:"sourceName,omitempty"`
}

// Returns the associated request.
func (o *GetSourceFilterListParams) GetRequestName() string {
	return "GetSourceFilterList"
}

// Represents the response body for the GetSourceFilterList request.
type GetSourceFilterListResponse struct {
	Filters []*typedefs.Filter `json:"filters,omitempty"`
}

// Gets an array of all of a source's filters.
func (c *Client) GetSourceFilterList(params *GetSourceFilterListParams) (*GetSourceFilterListResponse, error) {
	data := &GetSourceFilterListResponse{}
	return data, c.SendRequest(params, data)
}
