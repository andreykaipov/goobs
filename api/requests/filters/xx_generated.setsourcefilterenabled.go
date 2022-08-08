// This file has been automatically generated. Don't edit it.

package filters

// Represents the request body for the SetSourceFilterEnabled request.
type SetSourceFilterEnabledParams struct {
	// New enable state of the filter
	FilterEnabled *bool `json:"filterEnabled,omitempty"`

	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}

// Returns the associated request.
func (o *SetSourceFilterEnabledParams) GetRequestName() string {
	return "SetSourceFilterEnabled"
}

// Represents the response body for the SetSourceFilterEnabled request.
type SetSourceFilterEnabledResponse struct{}

// Sets the enable state of a source filter.
func (c *Client) SetSourceFilterEnabled(params *SetSourceFilterEnabledParams) (*SetSourceFilterEnabledResponse, error) {
	data := &SetSourceFilterEnabledResponse{}
	return data, c.SendRequest(params, data)
}
