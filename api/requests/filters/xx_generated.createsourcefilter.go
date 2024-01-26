// This file has been automatically generated. Don't edit it.

package filters

// Represents the request body for the CreateSourceFilter request.
type CreateSourceFilterParams struct {
	// The kind of filter to be created
	FilterKind *string `json:"filterKind,omitempty"`

	// Name of the new filter to be created
	FilterName *string `json:"filterName,omitempty"`

	// Settings object to initialize the filter with
	FilterSettings map[string]any `json:"filterSettings,omitempty"`

	// Name of the source to add the filter to
	SourceName *string `json:"sourceName,omitempty"`

	// UUID of the source to add the filter to
	SourceUuid *string `json:"sourceUuid,omitempty"`
}

func NewCreateSourceFilterParams() *CreateSourceFilterParams {
	return &CreateSourceFilterParams{}
}
func (o *CreateSourceFilterParams) WithFilterKind(x string) *CreateSourceFilterParams {
	o.FilterKind = &x
	return o
}
func (o *CreateSourceFilterParams) WithFilterName(x string) *CreateSourceFilterParams {
	o.FilterName = &x
	return o
}
func (o *CreateSourceFilterParams) WithFilterSettings(x map[string]any) *CreateSourceFilterParams {
	o.FilterSettings = x
	return o
}
func (o *CreateSourceFilterParams) WithSourceName(x string) *CreateSourceFilterParams {
	o.SourceName = &x
	return o
}
func (o *CreateSourceFilterParams) WithSourceUuid(x string) *CreateSourceFilterParams {
	o.SourceUuid = &x
	return o
}

// Returns the associated request.
func (o *CreateSourceFilterParams) GetRequestName() string {
	return "CreateSourceFilter"
}

// Represents the response body for the CreateSourceFilter request.
type CreateSourceFilterResponse struct {
	_response
}

// Creates a new filter, adding it to the specified source.
func (c *Client) CreateSourceFilter(params *CreateSourceFilterParams) (*CreateSourceFilterResponse, error) {
	data := &CreateSourceFilterResponse{}
	return data, c.client.SendRequest(params, data)
}
