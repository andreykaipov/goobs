// This file has been automatically generated. Don't edit it.

package sources

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
GetSourceFiltersParams represents the params body for the "GetSourceFilters" request.
List filters applied to a source
Since 4.5.0.
*/
type GetSourceFiltersParams struct {
	requests.ParamsBasic

	// Source name
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "GetSourceFilters".
func (o *GetSourceFiltersParams) GetSelfName() string {
	return "GetSourceFilters"
}

/*
GetSourceFiltersResponse represents the response body for the "GetSourceFilters" request.
List filters applied to a source
Since v4.5.0.
*/
type GetSourceFiltersResponse struct {
	requests.ResponseBasic

	// The filters for this object.
	Filters []*typedefs.Filter `json:"filters,omitempty"`
}

// GetSourceFilters sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceFilters(params *GetSourceFiltersParams) (*GetSourceFiltersResponse, error) {
	data := &GetSourceFiltersResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
