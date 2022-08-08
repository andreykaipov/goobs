// This file has been automatically generated. Don't edit it.

package filters

// Represents the request body for the RemoveSourceFilter request.
type RemoveSourceFilterParams struct {
	// Name of the filter to remove
	FilterName string `json:"filterName,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}

// Returns the associated request.
func (o *RemoveSourceFilterParams) GetRequestName() string {
	return "RemoveSourceFilter"
}

// Represents the response body for the RemoveSourceFilter request.
type RemoveSourceFilterResponse struct{}

// Removes a filter from a source.
func (c *Client) RemoveSourceFilter(params *RemoveSourceFilterParams) (*RemoveSourceFilterResponse, error) {
	data := &RemoveSourceFilterResponse{}
	return data, c.SendRequest(params, data)
}
