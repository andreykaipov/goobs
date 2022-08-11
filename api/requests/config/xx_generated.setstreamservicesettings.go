// This file has been automatically generated. Don't edit it.

package config

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the SetStreamServiceSettings request.
type SetStreamServiceSettingsParams struct {
	StreamServiceSettings *typedefs.StreamServiceSettings `json:"streamServiceSettings,omitempty"`

	// Type of stream service to apply. Example: `rtmp_common` or `rtmp_custom`
	StreamServiceType string `json:"streamServiceType,omitempty"`
}

// Returns the associated request.
func (o *SetStreamServiceSettingsParams) GetRequestName() string {
	return "SetStreamServiceSettings"
}

// Represents the response body for the SetStreamServiceSettings request.
type SetStreamServiceSettingsResponse struct{}

/*
Sets the current stream service settings (stream destination).

Note: Simple RTMP settings can be set with type `rtmp_custom` and the settings fields `server` and `key`.
*/
func (c *Client) SetStreamServiceSettings(
	params *SetStreamServiceSettingsParams,
) (*SetStreamServiceSettingsResponse, error) {
	data := &SetStreamServiceSettingsResponse{}
	return data, c.SendRequest(params, data)
}
