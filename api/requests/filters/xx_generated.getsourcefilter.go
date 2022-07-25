// This file has been automatically generated. Don't edit it.

package filters

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourceFilterParams represents the params body for the "GetSourceFilter" request.
Gets the info for a specific source filter.
*/
type GetSourceFilterParams struct {
	requests.ParamsBasic

	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// Name of the source
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "GetSourceFilter".
func (o *GetSourceFilterParams) GetSelfName() string {
	return "GetSourceFilter"
}

/*
GetSourceFilterResponse represents the response body for the "GetSourceFilter" request.
Gets the info for a specific source filter.
*/
type GetSourceFilterResponse struct {
	requests.ResponseBasic

	// Whether the filter is enabled
	FilterEnabled bool `json:"filterEnabled,omitempty"`

	// Index of the filter in the list, beginning at 0
	FilterIndex float64 `json:"filterIndex,omitempty"`

	// The kind of filter
	FilterKind string `json:"filterKind,omitempty"`

	// Settings object associated with the filter
	FilterSettings interface{} `json:"filterSettings,omitempty"`
}

// GetSourceFilter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceFilter(params *GetSourceFilterParams) (*GetSourceFilterResponse, error) {
	data := &GetSourceFilterResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
