// This file has been automatically generated. Don't edit it.

package filters

/*
CreateSourceFilterParams represents the params body for the "CreateSourceFilter" request.
Creates a new filter, adding it to the specified source.
*/
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

/*
CreateSourceFilterResponse represents the response body for the "CreateSourceFilter" request.
Creates a new filter, adding it to the specified source.
*/
type CreateSourceFilterResponse struct{}

// CreateSourceFilter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) CreateSourceFilter(params *CreateSourceFilterParams) (*CreateSourceFilterResponse, error) {
	resp, err := c.SendRequest("CreateSourceFilter", params)
	if err != nil {
		return nil, err
	}
	return resp.(*CreateSourceFilterResponse), nil
}
