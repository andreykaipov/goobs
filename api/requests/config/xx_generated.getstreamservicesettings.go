// This file has been automatically generated. Don't edit it.

package config

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetStreamServiceSettingsParams represents the params body for the "GetStreamServiceSettings" request.
Gets the current stream service settings (stream destination).
*/
type GetStreamServiceSettingsParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetStreamServiceSettings".
func (o *GetStreamServiceSettingsParams) GetSelfName() string {
	return "GetStreamServiceSettings"
}

/*
GetStreamServiceSettingsResponse represents the response body for the "GetStreamServiceSettings" request.
Gets the current stream service settings (stream destination).
*/
type GetStreamServiceSettingsResponse struct {
	requests.ResponseBasic

	// Stream service settings
	StreamServiceSettings interface{} `json:"streamServiceSettings,omitempty"`

	// Stream service type, like `rtmp_custom` or `rtmp_common`
	StreamServiceType string `json:"streamServiceType,omitempty"`
}

// GetStreamServiceSettings sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetStreamServiceSettings(
	paramss ...*GetStreamServiceSettingsParams,
) (*GetStreamServiceSettingsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStreamServiceSettingsParams{{}}
	}
	params := paramss[0]
	data := &GetStreamServiceSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
