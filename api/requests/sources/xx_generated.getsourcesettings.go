// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSourceSettingsParams represents the params body for the "GetSourceSettings" request.
Get settings of the specified source
Since 4.3.0.
*/
type GetSourceSettingsParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`

	// Type of the specified source. Useful for type-checking if you expect a specific settings schema.
	SourceType string `json:"sourceType,omitempty"`
}

// GetSelfName just returns "GetSourceSettings".
func (o *GetSourceSettingsParams) GetSelfName() string {
	return "GetSourceSettings"
}

/*
GetSourceSettingsResponse represents the response body for the "GetSourceSettings" request.
Get settings of the specified source
Since v4.3.0.
*/
type GetSourceSettingsResponse struct {
	requests.ResponseBasic

	// Source name
	SourceName string `json:"sourceName,omitempty"`

	// Source settings (varies between source types, may require some probing around).
	SourceSettings map[string]interface{} `json:"sourceSettings,omitempty"`

	// Type of the specified source
	SourceType string `json:"sourceType,omitempty"`
}

// GetSourceSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceSettings(params *GetSourceSettingsParams) (*GetSourceSettingsResponse, error) {
	data := &GetSourceSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
