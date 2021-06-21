// This file has been automatically generated. Don't edit it.

package streaming

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
SetStreamSettingsParams represents the params body for the "SetStreamSettings" request.
Sets one or more attributes of the current streaming server settings. Any options not passed will remain unchanged. Returns the updated settings in response. If 'type' is different than the current streaming service type, all settings are required. Returns the full settings of the stream (the same as GetStreamSettings).
Since 4.1.0.
*/
type SetStreamSettingsParams struct {
	requests.ParamsBasic

	// Persist the settings to disk.
	Save bool `json:"save"`

	//
	Settings *typedefs.StreamSettings `json:"settings,omitempty"`

	// The type of streaming service configuration, usually `rtmp_custom` or `rtmp_common`.
	Type string `json:"type,omitempty"`
}

// GetSelfName just returns "SetStreamSettings".
func (o *SetStreamSettingsParams) GetSelfName() string {
	return "SetStreamSettings"
}

/*
SetStreamSettingsResponse represents the response body for the "SetStreamSettings" request.
Sets one or more attributes of the current streaming server settings. Any options not passed will remain unchanged. Returns the updated settings in response. If 'type' is different than the current streaming service type, all settings are required. Returns the full settings of the stream (the same as GetStreamSettings).
Since v4.1.0.
*/
type SetStreamSettingsResponse struct {
	requests.ResponseBasic
}

// SetStreamSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetStreamSettings(params *SetStreamSettingsParams) (*SetStreamSettingsResponse, error) {
	data := &SetStreamSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
