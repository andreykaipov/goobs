// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
AddFilterToSourceParams represents the params body for the "AddFilterToSource" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#AddFilterToSource.
*/
type AddFilterToSourceParams struct {
	requests.ParamsBasic

	// Name of the new filter
	FilterName string `json:"filterName"`

	// Filter settings
	FilterSettings map[string]interface{} `json:"filterSettings"`

	// Filter type
	FilterType string `json:"filterType"`

	// Name of the source on which the filter is added
	SourceName string `json:"sourceName"`
}

// Name just returns "AddFilterToSource".
func (o *AddFilterToSourceParams) Name() string {
	return "AddFilterToSource"
}

/*
AddFilterToSourceResponse represents the response body for the "AddFilterToSource" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#AddFilterToSource.
*/
type AddFilterToSourceResponse struct {
	requests.ResponseBasic
}

// AddFilterToSource sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) AddFilterToSource(
	params *AddFilterToSourceParams,
) (*AddFilterToSourceResponse, error) {
	data := &AddFilterToSourceResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
