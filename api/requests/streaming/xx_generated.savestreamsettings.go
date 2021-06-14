// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SaveStreamSettingsParams represents the params body for the "SaveStreamSettings" request.
Save the current streaming server settings to disk.
Since 4.1.0.
*/
type SaveStreamSettingsParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "SaveStreamSettings".
func (o *SaveStreamSettingsParams) GetSelfName() string {
	return "SaveStreamSettings"
}

/*
SaveStreamSettingsResponse represents the response body for the "SaveStreamSettings" request.
Save the current streaming server settings to disk.
Since v4.1.0.
*/
type SaveStreamSettingsResponse struct {
	requests.ResponseBasic
}

// SaveStreamSettings sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) SaveStreamSettings(paramss ...*SaveStreamSettingsParams) (*SaveStreamSettingsResponse, error) {
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
