// This file has been automatically generated. Don't edit it.

package filters

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the GetSourceFilterList request.
type GetSourceFilterListParams struct {
	// Name of the source
	SourceName *string `json:"sourceName,omitempty"`

	// UUID of the source
	SourceUuid *string `json:"sourceUuid,omitempty"`
}

func NewGetSourceFilterListParams() *GetSourceFilterListParams {
	return &GetSourceFilterListParams{}
}
func (o *GetSourceFilterListParams) WithSourceName(x string) *GetSourceFilterListParams {
	o.SourceName = &x
	return o
}
func (o *GetSourceFilterListParams) WithSourceUuid(x string) *GetSourceFilterListParams {
	o.SourceUuid = &x
	return o
}

// Returns the associated request.
func (o *GetSourceFilterListParams) GetRequestName() string {
	return "GetSourceFilterList"
}

// Represents the response body for the GetSourceFilterList request.
type GetSourceFilterListResponse struct {
	_response

	// Array of filters
	Filters []*typedefs.Filter `json:"filters,omitempty"`
}

// Gets an array of all of a source's filters.
func (c *Client) GetSourceFilterList(paramss ...*GetSourceFilterListParams) (*GetSourceFilterListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSourceFilterListParams{{}}
	}
	params := paramss[0]
	data := &GetSourceFilterListResponse{}
	return data, c.client.SendRequest(params, data)
}
