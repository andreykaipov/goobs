// This file has been automatically generated. Don't edit it.

package filters

/*
GetSourceFilterParams represents the params body for the "GetSourceFilter" request.
Gets the info for a specific source filter.
*/
type GetSourceFilterParams struct {
	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// Name of the source
	SourceName string `json:"sourceName,omitempty"`
}

/*
GetSourceFilterResponse represents the response body for the "GetSourceFilter" request.
Gets the info for a specific source filter.
*/
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

// GetSourceFilter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceFilter(params *GetSourceFilterParams) (*GetSourceFilterResponse, error) {
	resp, err := c.SendRequest("GetSourceFilter", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetSourceFilterResponse), nil
}
