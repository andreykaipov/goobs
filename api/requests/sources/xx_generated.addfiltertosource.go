// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
AddFilterToSourceParams represents the params body for the "AddFilterToSource" request.
Add a new filter to a source. Available source types along with their settings properties are available from `GetSourceTypesList`.
Since 4.5.0.
*/
type AddFilterToSourceParams struct {
	requests.ParamsBasic

	// Name of the new filter
	FilterName string `json:"filterName,omitempty"`

	// Filter settings
	FilterSettings map[string]interface{} `json:"filterSettings,omitempty"`

	// Filter type
	FilterType string `json:"filterType,omitempty"`

	// Name of the source on which the filter is added
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "AddFilterToSource".
func (o *AddFilterToSourceParams) GetSelfName() string {
	return "AddFilterToSource"
}

/*
AddFilterToSourceResponse represents the response body for the "AddFilterToSource" request.
Add a new filter to a source. Available source types along with their settings properties are available from `GetSourceTypesList`.
Since v4.5.0.
*/
type AddFilterToSourceResponse struct {
	requests.ResponseBasic
}

// AddFilterToSource sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) AddFilterToSource(params *AddFilterToSourceParams) (*AddFilterToSourceResponse, error) {
	data := &AddFilterToSourceResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
