// This file has been automatically generated. Don't edit it.

package filters

import requests "github.com/andreykaipov/goobs/api/requests"

/*
RemoveSourceFilterParams represents the params body for the "RemoveSourceFilter" request.
Removes a filter from a source.
*/
type RemoveSourceFilterParams struct {
	requests.ParamsBasic

	// Name of the filter to remove
	FilterName string `json:"filterName,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "RemoveSourceFilter".
func (o *RemoveSourceFilterParams) GetSelfName() string {
	return "RemoveSourceFilter"
}

/*
RemoveSourceFilterResponse represents the response body for the "RemoveSourceFilter" request.
Removes a filter from a source.
*/
type RemoveSourceFilterResponse struct {
	requests.ResponseBasic
}

// RemoveSourceFilter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RemoveSourceFilter(params *RemoveSourceFilterParams) (*RemoveSourceFilterResponse, error) {
	data := &RemoveSourceFilterResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
