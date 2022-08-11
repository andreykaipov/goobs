// This file has been automatically generated. Don't edit it.

package config

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the GetStreamServiceSettings request.
type GetStreamServiceSettingsParams struct{}

// Returns the associated request.
func (o *GetStreamServiceSettingsParams) GetRequestName() string {
	return "GetStreamServiceSettings"
}

// Represents the response body for the GetStreamServiceSettings request.
type GetStreamServiceSettingsResponse struct {
	StreamServiceSettings *typedefs.StreamServiceSettings `json:"streamServiceSettings,omitempty"`

	// Stream service type, like `rtmp_custom` or `rtmp_common`
	StreamServiceType string `json:"streamServiceType,omitempty"`
}

// Gets the current stream service settings (stream destination).
func (c *Client) GetStreamServiceSettings(
	paramss ...*GetStreamServiceSettingsParams,
) (*GetStreamServiceSettingsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStreamServiceSettingsParams{{}}
	}
	params := paramss[0]
	data := &GetStreamServiceSettingsResponse{}
	return data, c.SendRequest(params, data)
}
