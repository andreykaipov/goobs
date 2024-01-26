// This file has been automatically generated. Don't edit it.

package filters

// Represents the request body for the GetSourceFilter request.
type GetSourceFilterParams struct {
	// Name of the filter
	FilterName *string `json:"filterName,omitempty"`

	// Name of the source
	SourceName *string `json:"sourceName,omitempty"`

	// UUID of the source
	SourceUuid *string `json:"sourceUuid,omitempty"`
}

func NewGetSourceFilterParams() *GetSourceFilterParams {
	return &GetSourceFilterParams{}
}
func (o *GetSourceFilterParams) WithFilterName(x string) *GetSourceFilterParams {
	o.FilterName = &x
	return o
}
func (o *GetSourceFilterParams) WithSourceName(x string) *GetSourceFilterParams {
	o.SourceName = &x
	return o
}
func (o *GetSourceFilterParams) WithSourceUuid(x string) *GetSourceFilterParams {
	o.SourceUuid = &x
	return o
}

// Returns the associated request.
func (o *GetSourceFilterParams) GetRequestName() string {
	return "GetSourceFilter"
}

// Represents the response body for the GetSourceFilter request.
type GetSourceFilterResponse struct {
	_response

	// Whether the filter is enabled
	FilterEnabled bool `json:"filterEnabled,omitempty"`

	// Index of the filter in the list, beginning at 0
	FilterIndex int `json:"filterIndex,omitempty"`

	// The kind of filter
	FilterKind string `json:"filterKind,omitempty"`

	// Settings object associated with the filter
	FilterSettings map[string]any `json:"filterSettings,omitempty"`
}

// Gets the info for a specific source filter.
func (c *Client) GetSourceFilter(params *GetSourceFilterParams) (*GetSourceFilterResponse, error) {
	data := &GetSourceFilterResponse{}
	return data, c.client.SendRequest(params, data)
}
