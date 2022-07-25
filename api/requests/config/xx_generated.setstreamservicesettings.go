// This file has been automatically generated. Don't edit it.

package config

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetStreamServiceSettingsParams represents the params body for the "SetStreamServiceSettings" request.
Sets the current stream service settings (stream destination).

Note: Simple RTMP settings can be set with type `rtmp_custom` and the settings fields `server` and `key`.
*/
type SetStreamServiceSettingsParams struct {
	requests.ParamsBasic

	// Settings to apply to the service
	StreamServiceSettings interface{} `json:"streamServiceSettings,omitempty"`

	// Type of stream service to apply. Example: `rtmp_common` or `rtmp_custom`
	StreamServiceType string `json:"streamServiceType,omitempty"`
}

// GetSelfName just returns "SetStreamServiceSettings".
func (o *SetStreamServiceSettingsParams) GetSelfName() string {
	return "SetStreamServiceSettings"
}

/*
SetStreamServiceSettingsResponse represents the response body for the "SetStreamServiceSettings" request.
Sets the current stream service settings (stream destination).

Note: Simple RTMP settings can be set with type `rtmp_custom` and the settings fields `server` and `key`.
*/
type SetStreamServiceSettingsResponse struct {
	requests.ResponseBasic
}

// SetStreamServiceSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetStreamServiceSettings(
	params *SetStreamServiceSettingsParams,
) (*SetStreamServiceSettingsResponse, error) {
	data := &SetStreamServiceSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
