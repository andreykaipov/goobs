// This file has been automatically generated. Don't edit it.

package filters

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSourceFilterNameParams represents the params body for the "SetSourceFilterName" request.
Sets the name of a source filter (rename).
*/
type SetSourceFilterNameParams struct {
	requests.ParamsBasic

	// Current name of the filter
	FilterName string `json:"filterName,omitempty"`

	// New name for the filter
	NewFilterName string `json:"newFilterName,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "SetSourceFilterName".
func (o *SetSourceFilterNameParams) GetSelfName() string {
	return "SetSourceFilterName"
}

/*
SetSourceFilterNameResponse represents the response body for the "SetSourceFilterName" request.
Sets the name of a source filter (rename).
*/
type SetSourceFilterNameResponse struct {
	requests.ResponseBasic
}

// SetSourceFilterName sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceFilterName(params *SetSourceFilterNameParams) (*SetSourceFilterNameResponse, error) {
	data := &SetSourceFilterNameResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
