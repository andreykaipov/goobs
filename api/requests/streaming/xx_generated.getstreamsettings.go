// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetStreamSettingsParams represents the params body for the "GetStreamSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStreamSettings.
*/
type GetStreamSettingsParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetStreamSettingsParams
func (o *GetStreamSettingsParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetStreamSettingsParams
func (o *GetStreamSettingsParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetStreamSettingsParams
func (o *GetStreamSettingsParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetStreamSettingsResponse represents the response body for the "GetStreamSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStreamSettings.
*/
type GetStreamSettingsResponse struct {
	requests.Response

	Settings struct {
		// The publish key of the stream.
		Key string `json:"key"`

		// The password to use when accessing the streaming server. Only present if `use-auth` is
		// `true`.
		Password string `json:"password"`

		// The publish URL.
		Server string `json:"server"`

		// Indicates whether authentication should be used when connecting to the streaming server.
		UseAuth bool `json:"use-auth"`

		// The username to use when accessing the streaming server. Only present if `use-auth` is
		// `true`.
		Username string `json:"username"`
	} `json:"settings"`

	// The type of streaming service configuration. Possible values: 'rtmp_custom' or 'rtmp_common'.
	Type string `json:"type"`
}

// GetMessageID returns the MessageID of GetStreamSettingsResponse
func (o *GetStreamSettingsResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetStreamSettingsResponse
func (o *GetStreamSettingsResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetStreamSettingsResponse
func (o *GetStreamSettingsResponse) GetError() string {
	return o.Error
}

// GetStreamSettings sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetStreamSettings(
	paramss ...*GetStreamSettingsParams,
) (*GetStreamSettingsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStreamSettingsParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetStreamSettings"
	data := &GetStreamSettingsResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
