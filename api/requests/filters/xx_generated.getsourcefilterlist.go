// This file has been automatically generated. Don't edit it.

package filters

import (
	api "github.com/andreykaipov/goobs/api"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

// Represents the request body for the GetSourceFilterList request.
type GetSourceFilterListParams struct {
	// Name of the source
	SourceName *string `json:"sourceName,omitempty"`
}

func NewGetSourceFilterListParams() *GetSourceFilterListParams {
	return &GetSourceFilterListParams{}
}
func (o *GetSourceFilterListParams) WithSourceName(x string) *GetSourceFilterListParams {
	o.SourceName = &x
	return o
}

// Returns the associated request.
func (o *GetSourceFilterListParams) GetRequestName() string {
	return "GetSourceFilterList"
}

// Represents the response body for the GetSourceFilterList request.
type GetSourceFilterListResponse struct {
	api.ResponseCommon

	// Array of filters
	Filters []*typedefs.Filter `json:"filters,omitempty"`
}

// Gets an array of all of a source's filters.
func (c *Client) GetSourceFilterList(params *GetSourceFilterListParams) (*GetSourceFilterListResponse, error) {
	data := &GetSourceFilterListResponse{}
	return data, c.SendRequest(params, data)
}
