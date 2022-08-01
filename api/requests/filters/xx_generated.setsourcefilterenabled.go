// This file has been automatically generated. Don't edit it.

package filters

/*
SetSourceFilterEnabledParams represents the params body for the "SetSourceFilterEnabled" request.
Sets the enable state of a source filter.
*/
type SetSourceFilterEnabledParams struct {
	// New enable state of the filter
	FilterEnabled bool `json:"filterEnabled,omitempty"`

	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}

/*
SetSourceFilterEnabledResponse represents the response body for the "SetSourceFilterEnabled" request.
Sets the enable state of a source filter.
*/
type SetSourceFilterEnabledResponse struct{}

// SetSourceFilterEnabled sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceFilterEnabled(params *SetSourceFilterEnabledParams) (*SetSourceFilterEnabledResponse, error) {
	resp, err := c.SendRequest("SetSourceFilterEnabled", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetSourceFilterEnabledResponse), nil
}
