// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourceDefaultSettingsParams represents the params body for the "GetSourceDefaultSettings" request.
Get the default settings for a given source type.
Since 4.9.0.
*/
type GetSourceDefaultSettingsParams struct {
	requests.ParamsBasic

	// Source kind. Also called "source id" in libobs terminology.
	SourceKind string `json:"sourceKind,omitempty"`
}

// GetSelfName just returns "GetSourceDefaultSettings".
func (o *GetSourceDefaultSettingsParams) GetSelfName() string {
	return "GetSourceDefaultSettings"
}

/*
GetSourceDefaultSettingsResponse represents the response body for the "GetSourceDefaultSettings" request.
Get the default settings for a given source type.
Since v4.9.0.
*/
type GetSourceDefaultSettingsResponse struct {
	requests.ResponseBasic

	// Settings object for source.
	DefaultSettings map[string]interface{} `json:"defaultSettings,omitempty"`

	// Source kind. Same value as the `sourceKind` parameter.
	SourceKind string `json:"sourceKind,omitempty"`
}

// GetSourceDefaultSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceDefaultSettings(
	params *GetSourceDefaultSettingsParams,
) (*GetSourceDefaultSettingsResponse, error) {
	data := &GetSourceDefaultSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
