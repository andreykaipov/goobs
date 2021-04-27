// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SaveStreamSettingsParams represents the params body for the "SaveStreamSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SaveStreamSettings.
*/
type SaveStreamSettingsParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of SaveStreamSettingsParams
func (o *SaveStreamSettingsParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SaveStreamSettingsParams
func (o *SaveStreamSettingsParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SaveStreamSettingsParams
func (o *SaveStreamSettingsParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SaveStreamSettingsResponse represents the response body for the "SaveStreamSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SaveStreamSettings.
*/
type SaveStreamSettingsResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SaveStreamSettingsResponse
func (o *SaveStreamSettingsResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SaveStreamSettingsResponse
func (o *SaveStreamSettingsResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SaveStreamSettingsResponse
func (o *SaveStreamSettingsResponse) GetError() string {
	return o.Error
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
	params.RequestType = "SaveStreamSettings"
	data := &SaveStreamSettingsResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
