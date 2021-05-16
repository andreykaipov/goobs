// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSourceFilterSettingsParams represents the params body for the "SetSourceFilterSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceFilterSettings.
*/
type SetSourceFilterSettingsParams struct {
	requests.ParamsBasic

	// Name of the filter to reconfigure
	FilterName string `json:"filterName"`

	// New settings. These will be merged to the current filter settings.
	FilterSettings map[string]interface{} `json:"filterSettings"`

	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}

// Name just returns "SetSourceFilterSettings".
func (o *SetSourceFilterSettingsParams) Name() string {
	return "SetSourceFilterSettings"
}

/*
SetSourceFilterSettingsResponse represents the response body for the "SetSourceFilterSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSourceFilterSettings.
*/
type SetSourceFilterSettingsResponse struct {
	requests.ResponseBasic
}

// SetSourceFilterSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceFilterSettings(
	params *SetSourceFilterSettingsParams,
) (*SetSourceFilterSettingsResponse, error) {
	data := &SetSourceFilterSettingsResponse{}
	if err := c.WriteMessage(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
