// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourceFilterInfoParams represents the params body for the "GetSourceFilterInfo" request.
List filters applied to a source
Since 4.7.0.
*/
type GetSourceFilterInfoParams struct {
	requests.ParamsBasic

	// Source filter name
	FilterName string `json:"filterName,omitempty"`

	// Source name
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "GetSourceFilterInfo".
func (o *GetSourceFilterInfoParams) GetSelfName() string {
	return "GetSourceFilterInfo"
}

/*
GetSourceFilterInfoResponse represents the response body for the "GetSourceFilterInfo" request.
List filters applied to a source
Since v4.7.0.
*/
type GetSourceFilterInfoResponse struct {
	requests.ResponseBasic

	// Filter status (enabled or not)
	Enabled bool `json:"enabled"`

	// Filter name
	Name string `json:"name,omitempty"`

	// Filter settings
	Settings map[string]interface{} `json:"settings,omitempty"`

	// Filter type
	Type string `json:"type,omitempty"`
}

// GetSourceFilterInfo sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceFilterInfo(params *GetSourceFilterInfoParams) (*GetSourceFilterInfoResponse, error) {
	data := &GetSourceFilterInfoResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
