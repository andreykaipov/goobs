// This file has been automatically generated. Don't edit it.

package filters

// Represents the request body for the SetSourceFilterName request.
type SetSourceFilterNameParams struct {
	// Current name of the filter
	FilterName *string `json:"filterName,omitempty"`

	// New name for the filter
	NewFilterName *string `json:"newFilterName,omitempty"`

	// Name of the source the filter is on
	SourceName *string `json:"sourceName,omitempty"`

	// UUID of the source the filter is on
	SourceUuid *string `json:"sourceUuid,omitempty"`
}

func NewSetSourceFilterNameParams() *SetSourceFilterNameParams {
	return &SetSourceFilterNameParams{}
}
func (o *SetSourceFilterNameParams) WithFilterName(x string) *SetSourceFilterNameParams {
	o.FilterName = &x
	return o
}
func (o *SetSourceFilterNameParams) WithNewFilterName(x string) *SetSourceFilterNameParams {
	o.NewFilterName = &x
	return o
}
func (o *SetSourceFilterNameParams) WithSourceName(x string) *SetSourceFilterNameParams {
	o.SourceName = &x
	return o
}
func (o *SetSourceFilterNameParams) WithSourceUuid(x string) *SetSourceFilterNameParams {
	o.SourceUuid = &x
	return o
}

// Returns the associated request.
func (o *SetSourceFilterNameParams) GetRequestName() string {
	return "SetSourceFilterName"
}

// Represents the response body for the SetSourceFilterName request.
type SetSourceFilterNameResponse struct {
	_response
}

// Sets the name of a source filter (rename).
func (c *Client) SetSourceFilterName(params *SetSourceFilterNameParams) (*SetSourceFilterNameResponse, error) {
	data := &SetSourceFilterNameResponse{}
	return data, c.client.SendRequest(params, data)
}
