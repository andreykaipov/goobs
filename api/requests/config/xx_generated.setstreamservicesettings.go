// This file has been automatically generated. Don't edit it.

package config

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
SetStreamServiceSettingsParams represents the params body for the "SetStreamServiceSettings" request.
Sets the current stream service settings (stream destination).

Note: Simple RTMP settings can be set with type `rtmp_custom` and the settings fields `server` and `key`.
*/
type SetStreamServiceSettingsParams struct {
	//
	StreamServiceSettings *typedefs.StreamServiceSettings `json:"streamServiceSettings,omitempty"`

	// Type of stream service to apply. Example: `rtmp_common` or `rtmp_custom`
	StreamServiceType string `json:"streamServiceType,omitempty"`
}

/*
SetStreamServiceSettingsResponse represents the response body for the "SetStreamServiceSettings" request.
Sets the current stream service settings (stream destination).

Note: Simple RTMP settings can be set with type `rtmp_custom` and the settings fields `server` and `key`.
*/
type SetStreamServiceSettingsResponse struct{}

// SetStreamServiceSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetStreamServiceSettings(
	params *SetStreamServiceSettingsParams,
) (*SetStreamServiceSettingsResponse, error) {
	resp, err := c.SendRequest("SetStreamServiceSettings", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetStreamServiceSettingsResponse), nil
}
