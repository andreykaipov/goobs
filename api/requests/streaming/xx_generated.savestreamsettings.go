// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SaveStreamSettingsParams represents the params body for the "SaveStreamSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SaveStreamSettings.
*/
type SaveStreamSettingsParams struct {
	requests.ParamsBasic
}

// Name just returns "SaveStreamSettings".
func (o *SaveStreamSettingsParams) Name() string {
	return "SaveStreamSettings"
}

/*
SaveStreamSettingsResponse represents the response body for the "SaveStreamSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SaveStreamSettings.
*/
type SaveStreamSettingsResponse struct {
	requests.ResponseBasic
}

// SaveStreamSettings sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) SaveStreamSettings(
	paramss ...*SaveStreamSettingsParams,
) (*SaveStreamSettingsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*SaveStreamSettingsParams{{}}
	}
	params := paramss[0]
	data := &SaveStreamSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
