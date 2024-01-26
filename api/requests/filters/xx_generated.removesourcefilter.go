// This file has been automatically generated. Don't edit it.

package filters

// Represents the request body for the RemoveSourceFilter request.
type RemoveSourceFilterParams struct {
	// Name of the filter to remove
	FilterName *string `json:"filterName,omitempty"`

	// Name of the source the filter is on
	SourceName *string `json:"sourceName,omitempty"`

	// UUID of the source the filter is on
	SourceUuid *string `json:"sourceUuid,omitempty"`
}

func NewRemoveSourceFilterParams() *RemoveSourceFilterParams {
	return &RemoveSourceFilterParams{}
}
func (o *RemoveSourceFilterParams) WithFilterName(x string) *RemoveSourceFilterParams {
	o.FilterName = &x
	return o
}
func (o *RemoveSourceFilterParams) WithSourceName(x string) *RemoveSourceFilterParams {
	o.SourceName = &x
	return o
}
func (o *RemoveSourceFilterParams) WithSourceUuid(x string) *RemoveSourceFilterParams {
	o.SourceUuid = &x
	return o
}

// Returns the associated request.
func (o *RemoveSourceFilterParams) GetRequestName() string {
	return "RemoveSourceFilter"
}

// Represents the response body for the RemoveSourceFilter request.
type RemoveSourceFilterResponse struct {
	_response
}

// Removes a filter from a source.
func (c *Client) RemoveSourceFilter(params *RemoveSourceFilterParams) (*RemoveSourceFilterResponse, error) {
	data := &RemoveSourceFilterResponse{}
	return data, c.client.SendRequest(params, data)
}
