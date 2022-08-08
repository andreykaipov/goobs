// This file has been automatically generated. Don't edit it.

package filters

// Represents the request body for the GetSourceFilter request.
type GetSourceFilterParams struct {
	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// Name of the source
	SourceName string `json:"sourceName,omitempty"`
}

// Returns the associated request.
func (o *GetSourceFilterParams) GetRequestName() string {
	return "GetSourceFilter"
}

// Represents the response body for the GetSourceFilter request.
type GetSourceFilterResponse struct {
	// Whether the filter is enabled
	FilterEnabled bool `json:"filterEnabled,omitempty"`

	// Index of the filter in the list, beginning at 0
	FilterIndex float64 `json:"filterIndex,omitempty"`

	// The kind of filter
	FilterKind string `json:"filterKind,omitempty"`

	// Settings object associated with the filter
	FilterSettings map[string]interface{} `json:"filterSettings,omitempty"`
}

// Gets the info for a specific source filter.
func (c *Client) GetSourceFilter(params *GetSourceFilterParams) (*GetSourceFilterResponse, error) {
	data := &GetSourceFilterResponse{}
	return data, c.SendRequest(params, data)
}
