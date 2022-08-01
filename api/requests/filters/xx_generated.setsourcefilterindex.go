// This file has been automatically generated. Don't edit it.

package filters

/*
SetSourceFilterIndexParams represents the params body for the "SetSourceFilterIndex" request.
Sets the index position of a filter on a source.
*/
type SetSourceFilterIndexParams struct {
	// New index position of the filter
	FilterIndex float64 `json:"filterIndex,omitempty"`

	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}

/*
SetSourceFilterIndexResponse represents the response body for the "SetSourceFilterIndex" request.
Sets the index position of a filter on a source.
*/
type SetSourceFilterIndexResponse struct{}

// SetSourceFilterIndex sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceFilterIndex(params *SetSourceFilterIndexParams) (*SetSourceFilterIndexResponse, error) {
	resp, err := c.SendRequest("SetSourceFilterIndex", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetSourceFilterIndexResponse), nil
}
