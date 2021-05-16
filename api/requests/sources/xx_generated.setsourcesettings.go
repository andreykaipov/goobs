// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSourceSettingsParams represents the params body for the "SetSourceSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceSettings.
*/
type SetSourceSettingsParams struct {
	requests.ParamsBasic

	// Source name.
	SourceName string `json:"sourceName"`

	// Source settings (varies between source types, may require some probing around).
	SourceSettings map[string]interface{} `json:"sourceSettings"`

	// Type of the specified source. Useful for type-checking to avoid settings a set of settings
	// incompatible with the actual source's type.
	SourceType string `json:"sourceType"`
}

// Name just returns "SetSourceSettings".
func (o *SetSourceSettingsParams) Name() string {
	return "SetSourceSettings"
}

/*
SetSourceSettingsResponse represents the response body for the "SetSourceSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceSettings.
*/
type SetSourceSettingsResponse struct {
	requests.ResponseBasic

	// Source name
	SourceName string `json:"sourceName"`

	// Updated source settings
	SourceSettings map[string]interface{} `json:"sourceSettings"`

	// Type of the specified source
	SourceType string `json:"sourceType"`
}

// SetSourceSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceSettings(
	params *SetSourceSettingsParams,
) (*SetSourceSettingsResponse, error) {
	data := &SetSourceSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
