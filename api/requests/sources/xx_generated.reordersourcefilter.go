// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ReorderSourceFilterParams represents the params body for the "ReorderSourceFilter" request.
Move a filter in the chain (absolute index positioning)
Since 4.5.0.
*/
type ReorderSourceFilterParams struct {
	requests.ParamsBasic

	// Name of the filter to reorder
	FilterName string `json:"filterName,omitempty"`

	// Desired position of the filter in the chain
	NewIndex int `json:"newIndex,omitempty"`

	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "ReorderSourceFilter".
func (o *ReorderSourceFilterParams) GetSelfName() string {
	return "ReorderSourceFilter"
}

/*
ReorderSourceFilterResponse represents the response body for the "ReorderSourceFilter" request.
Move a filter in the chain (absolute index positioning)
Since v4.5.0.
*/
type ReorderSourceFilterResponse struct {
	requests.ResponseBasic
}

// ReorderSourceFilter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ReorderSourceFilter(params *ReorderSourceFilterParams) (*ReorderSourceFilterResponse, error) {
	data := &ReorderSourceFilterResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
