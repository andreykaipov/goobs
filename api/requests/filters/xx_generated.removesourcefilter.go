// This file has been automatically generated. Don't edit it.

package filters

/*
RemoveSourceFilterParams represents the params body for the "RemoveSourceFilter" request.
Removes a filter from a source.
*/
type RemoveSourceFilterParams struct {
	// Name of the filter to remove
	FilterName string `json:"filterName,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}

/*
RemoveSourceFilterResponse represents the response body for the "RemoveSourceFilter" request.
Removes a filter from a source.
*/
type RemoveSourceFilterResponse struct{}

// RemoveSourceFilter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RemoveSourceFilter(params *RemoveSourceFilterParams) (*RemoveSourceFilterResponse, error) {
	resp, err := c.SendRequest("RemoveSourceFilter", params)
	if err != nil {
		return nil, err
	}
	return resp.(*RemoveSourceFilterResponse), nil
}
