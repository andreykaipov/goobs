// This file has been automatically generated. Don't edit it.

package filters

// Represents the request body for the CreateSourceFilter request.
type CreateSourceFilterParams struct {
	// The kind of filter to be created
	FilterKind string `json:"filterKind,omitempty"`

	// Name of the new filter to be created
	FilterName string `json:"filterName,omitempty"`

	// Settings object to initialize the filter with
	FilterSettings map[string]interface{} `json:"filterSettings,omitempty"`

	// Name of the source to add the filter to
	SourceName string `json:"sourceName,omitempty"`
}

// Returns the associated request.
func (o *CreateSourceFilterParams) GetRequestName() string {
	return "CreateSourceFilter"
}

// Represents the response body for the CreateSourceFilter request.
type CreateSourceFilterResponse struct{}

// Creates a new filter, adding it to the specified source.
func (c *Client) CreateSourceFilter(params *CreateSourceFilterParams) (*CreateSourceFilterResponse, error) {
	data := &CreateSourceFilterResponse{}
	return data, c.SendRequest(params, data)
}
