// This file has been automatically generated. Don't edit it.

package filters

// Represents the request body for the SetSourceFilterIndex request.
type SetSourceFilterIndexParams struct {
	// New index position of the filter
	FilterIndex float64 `json:"filterIndex,omitempty"`

	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}

func NewSetSourceFilterIndexParams() *SetSourceFilterIndexParams {
	return &SetSourceFilterIndexParams{}
}
func (o *SetSourceFilterIndexParams) WithFilterIndex(x float64) *SetSourceFilterIndexParams {
	o.FilterIndex = x
	return o
}
func (o *SetSourceFilterIndexParams) WithFilterName(x string) *SetSourceFilterIndexParams {
	o.FilterName = x
	return o
}
func (o *SetSourceFilterIndexParams) WithSourceName(x string) *SetSourceFilterIndexParams {
	o.SourceName = x
	return o
}

// Returns the associated request.
func (o *SetSourceFilterIndexParams) GetRequestName() string {
	return "SetSourceFilterIndex"
}

// Represents the response body for the SetSourceFilterIndex request.
type SetSourceFilterIndexResponse struct{}

// Sets the index position of a filter on a source.
func (c *Client) SetSourceFilterIndex(params *SetSourceFilterIndexParams) (*SetSourceFilterIndexResponse, error) {
	data := &SetSourceFilterIndexResponse{}
	return data, c.SendRequest(params, data)
}
