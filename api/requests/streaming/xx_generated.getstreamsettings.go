// This file has been automatically generated. Don't edit it.

package streaming

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
GetStreamSettingsParams represents the params body for the "GetStreamSettings" request.
Get the current streaming server settings.
Since 4.1.0.
*/
type GetStreamSettingsParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetStreamSettings".
func (o *GetStreamSettingsParams) GetSelfName() string {
	return "GetStreamSettings"
}

/*
GetStreamSettingsResponse represents the response body for the "GetStreamSettings" request.
Get the current streaming server settings.
Since v4.1.0.
*/
type GetStreamSettingsResponse struct {
	requests.ResponseBasic

	//
	Settings *typedefs.StreamSettings `json:"settings,omitempty"`

	// The type of streaming service configuration. Possible values: 'rtmp_custom' or 'rtmp_common'.
	Type string `json:"type,omitempty"`
}

// GetStreamSettings sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) GetStreamSettings(paramss ...*GetStreamSettingsParams) (*GetStreamSettingsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStreamSettingsParams{{}}
	}
	params := paramss[0]
	data := &GetStreamSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
