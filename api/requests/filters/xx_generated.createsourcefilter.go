// This file has been automatically generated. Don't edit it.

package filters

import requests "github.com/andreykaipov/goobs/api/requests"

/*
CreateSourceFilterParams represents the params body for the "CreateSourceFilter" request.
Creates a new filter, adding it to the specified source.
*/
type CreateSourceFilterParams struct {
	requests.ParamsBasic

	// The kind of filter to be created
	FilterKind string `json:"filterKind,omitempty"`

	// Name of the new filter to be created
	FilterName string `json:"filterName,omitempty"`

	// Settings object to initialize the filter with
	FilterSettings interface{} `json:"filterSettings,omitempty"`

	// Name of the source to add the filter to
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "CreateSourceFilter".
func (o *CreateSourceFilterParams) GetSelfName() string {
	return "CreateSourceFilter"
}

/*
CreateSourceFilterResponse represents the response body for the "CreateSourceFilter" request.
Creates a new filter, adding it to the specified source.
*/
type CreateSourceFilterResponse struct {
	requests.ResponseBasic
}

// CreateSourceFilter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) CreateSourceFilter(params *CreateSourceFilterParams) (*CreateSourceFilterResponse, error) {
	data := &CreateSourceFilterResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
