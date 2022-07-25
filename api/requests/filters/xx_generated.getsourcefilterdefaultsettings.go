// This file has been automatically generated. Don't edit it.

package filters

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourceFilterDefaultSettingsParams represents the params body for the "GetSourceFilterDefaultSettings" request.
Gets the default settings for a filter kind.
*/
type GetSourceFilterDefaultSettingsParams struct {
	requests.ParamsBasic

	// Filter kind to get the default settings for
	FilterKind string `json:"filterKind,omitempty"`
}

// GetSelfName just returns "GetSourceFilterDefaultSettings".
func (o *GetSourceFilterDefaultSettingsParams) GetSelfName() string {
	return "GetSourceFilterDefaultSettings"
}

/*
GetSourceFilterDefaultSettingsResponse represents the response body for the "GetSourceFilterDefaultSettings" request.
Gets the default settings for a filter kind.
*/
type GetSourceFilterDefaultSettingsResponse struct {
	requests.ResponseBasic

	// Object of default settings for the filter kind
	DefaultFilterSettings interface{} `json:"defaultFilterSettings,omitempty"`
}

// GetSourceFilterDefaultSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceFilterDefaultSettings(
	params *GetSourceFilterDefaultSettingsParams,
) (*GetSourceFilterDefaultSettingsResponse, error) {
	data := &GetSourceFilterDefaultSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
