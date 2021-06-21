// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSourceSettingsParams represents the params body for the "SetSourceSettings" request.
Set settings of the specified source.
Since 4.3.0.
*/
type SetSourceSettingsParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName,omitempty"`

	// Source settings (varies between source types, may require some probing around).
	SourceSettings map[string]interface{} `json:"sourceSettings,omitempty"`

	// Type of the specified source. Useful for type-checking to avoid settings a set of settings incompatible with the
	// actual source's type.
	SourceType string `json:"sourceType,omitempty"`
}

// GetSelfName just returns "SetSourceSettings".
func (o *SetSourceSettingsParams) GetSelfName() string {
	return "SetSourceSettings"
}

/*
SetSourceSettingsResponse represents the response body for the "SetSourceSettings" request.
Set settings of the specified source.
Since v4.3.0.
*/
type SetSourceSettingsResponse struct {
	requests.ResponseBasic

	// Source name
	SourceName string `json:"sourceName,omitempty"`

	// Updated source settings
	SourceSettings map[string]interface{} `json:"sourceSettings,omitempty"`

	// Type of the specified source
	SourceType string `json:"sourceType,omitempty"`
}

// SetSourceSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceSettings(params *SetSourceSettingsParams) (*SetSourceSettingsResponse, error) {
	data := &SetSourceSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
