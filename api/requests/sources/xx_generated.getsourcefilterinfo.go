// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourceFilterInfoParams represents the params body for the "GetSourceFilterInfo" request.
List filters applied to a source

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#GetSourceFilterInfo.
*/
type GetSourceFilterInfoParams struct {
	requests.ParamsBasic

	// Source filter name
	FilterName string `json:"filterName"`

	// Source name
	SourceName string `json:"sourceName"`
}

// Name just returns "GetSourceFilterInfo".
func (o *GetSourceFilterInfoParams) Name() string {
	return "GetSourceFilterInfo"
}

/*
GetSourceFilterInfoResponse represents the response body for the "GetSourceFilterInfo" request.
List filters applied to a source

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#GetSourceFilterInfo.
*/
type GetSourceFilterInfoResponse struct {
	requests.ResponseBasic

	// Filter status (enabled or not)
	Enabled bool `json:"enabled"`

	// Filter name
	Name string `json:"name"`

	// Filter settings
	Settings map[string]interface{} `json:"settings"`

	// Filter type
	Type string `json:"type"`
}

// GetSourceFilterInfo sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceFilterInfo(
	params *GetSourceFilterInfoParams,
) (*GetSourceFilterInfoResponse, error) {
	data := &GetSourceFilterInfoResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
