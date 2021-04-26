// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetStreamingStatusParams represents the params body for the "GetStreamingStatus" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStreamingStatus.
*/
type GetStreamingStatusParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetStreamingStatusParams
func (o *GetStreamingStatusParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetStreamingStatusParams
func (o *GetStreamingStatusParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetStreamingStatusParams
func (o *GetStreamingStatusParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetStreamingStatusResponse represents the response body for the "GetStreamingStatus" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStreamingStatus.
*/
type GetStreamingStatusResponse struct {
	requests.Response

	// Always false. Retrocompatibility with OBSRemote.
	PreviewOnly bool `json:"preview-only"`

	// Time elapsed since recording started (only present if currently recording).
	RecTimecode string `json:"rec-timecode"`

	// Current recording status.
	Recording bool `json:"recording"`

	// Time elapsed since streaming started (only present if currently streaming).
	StreamTimecode string `json:"stream-timecode"`

	// Current streaming status.
	Streaming bool `json:"streaming"`
}

// GetMessageID returns the MessageID of GetStreamingStatusResponse
func (o *GetStreamingStatusResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetStreamingStatusResponse
func (o *GetStreamingStatusResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetStreamingStatusResponse
func (o *GetStreamingStatusResponse) GetError() string {
	return o.Error
}

// GetStreamingStatus sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetStreamingStatus(
	paramss ...*GetStreamingStatusParams,
) (*GetStreamingStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStreamingStatusParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetStreamingStatus"
	data := &GetStreamingStatusResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
StartStopStreamingParams represents the params body for the "StartStopStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopStreaming.
*/
type StartStopStreamingParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of StartStopStreamingParams
func (o *StartStopStreamingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StartStopStreamingParams
func (o *StartStopStreamingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StartStopStreamingParams
func (o *StartStopStreamingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StartStopStreamingResponse represents the response body for the "StartStopStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopStreaming.
*/
type StartStopStreamingResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StartStopStreamingResponse
func (o *StartStopStreamingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StartStopStreamingResponse
func (o *StartStopStreamingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StartStopStreamingResponse
func (o *StartStopStreamingResponse) GetError() string {
	return o.Error
}

// StartStopStreaming sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) StartStopStreaming(
	paramss ...*StartStopStreamingParams,
) (*StartStopStreamingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartStopStreamingParams{{}}
	}
	params := paramss[0]
	params.RequestType = "StartStopStreaming"
	data := &StartStopStreamingResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
StartStreamingParams represents the params body for the "StartStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStreaming.
*/
type StartStreamingParams struct {
	requests.Params

	Stream struct {
		// Adds the given object parameters as encoded query string parameters to the 'key' of the
		// RTMP stream. Used to pass data to the RTMP service about the streaming. May be any
		// String, Numeric, or Boolean field.
		Metadata map[string]interface{} `json:"metadata"`

		Settings struct {
			// The publish key of the stream.
			Key string `json:"key"`

			// If authentication is enabled, the password for the streaming server. Ignored if
			// `use-auth` is not set to `true`.
			Password string `json:"password"`

			// The publish URL.
			Server string `json:"server"`

			// Indicates whether authentication should be used when connecting to the streaming
			// server.
			UseAuth bool `json:"use-auth"`

			// If authentication is enabled, the username for the streaming server. Ignored if
			// `use-auth` is not set to `true`.
			Username string `json:"username"`
		} `json:"settings"`

		// If specified ensures the type of stream matches the given type (usually 'rtmp_custom' or
		// 'rtmp_common'). If the currently configured stream type does not match the given stream
		// type, all settings must be specified in the `settings` object or an error will occur when
		// starting the stream.
		Type string `json:"type"`
	} `json:"stream"`
}

// GetRequestType returns the RequestType of StartStreamingParams
func (o *StartStreamingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StartStreamingParams
func (o *StartStreamingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StartStreamingParams
func (o *StartStreamingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StartStreamingResponse represents the response body for the "StartStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStreaming.
*/
type StartStreamingResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StartStreamingResponse
func (o *StartStreamingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StartStreamingResponse
func (o *StartStreamingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StartStreamingResponse
func (o *StartStreamingResponse) GetError() string {
	return o.Error
}

// StartStreaming sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) StartStreaming(params *StartStreamingParams) (*StartStreamingResponse, error) {
	params.RequestType = "StartStreaming"
	data := &StartStreamingResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
StopStreamingParams represents the params body for the "StopStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopStreaming.
*/
type StopStreamingParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of StopStreamingParams
func (o *StopStreamingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StopStreamingParams
func (o *StopStreamingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StopStreamingParams
func (o *StopStreamingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StopStreamingResponse represents the response body for the "StopStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopStreaming.
*/
type StopStreamingResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StopStreamingResponse
func (o *StopStreamingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StopStreamingResponse
func (o *StopStreamingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StopStreamingResponse
func (o *StopStreamingResponse) GetError() string {
	return o.Error
}

// StopStreaming sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) StopStreaming(paramss ...*StopStreamingParams) (*StopStreamingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopStreamingParams{{}}
	}
	params := paramss[0]
	params.RequestType = "StopStreaming"
	data := &StopStreamingResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
SetStreamSettingsParams represents the params body for the "SetStreamSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetStreamSettings.
*/
type SetStreamSettingsParams struct {
	requests.Params

	// Persist the settings to disk.
	Save bool `json:"save"`

	Settings struct {
		// The publish key.
		Key string `json:"key"`

		// The password for the streaming service.
		Password string `json:"password"`

		// The publish URL.
		Server string `json:"server"`

		// Indicates whether authentication should be used when connecting to the streaming server.
		UseAuth bool `json:"use-auth"`

		// The username for the streaming service.
		Username string `json:"username"`
	} `json:"settings"`

	// The type of streaming service configuration, usually `rtmp_custom` or `rtmp_common`.
	Type string `json:"type"`
}

// GetRequestType returns the RequestType of SetStreamSettingsParams
func (o *SetStreamSettingsParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetStreamSettingsParams
func (o *SetStreamSettingsParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetStreamSettingsParams
func (o *SetStreamSettingsParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetStreamSettingsResponse represents the response body for the "SetStreamSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetStreamSettings.
*/
type SetStreamSettingsResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetStreamSettingsResponse
func (o *SetStreamSettingsResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetStreamSettingsResponse
func (o *SetStreamSettingsResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetStreamSettingsResponse
func (o *SetStreamSettingsResponse) GetError() string {
	return o.Error
}

// SetStreamSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetStreamSettings(
	params *SetStreamSettingsParams,
) (*SetStreamSettingsResponse, error) {
	params.RequestType = "SetStreamSettings"
	data := &SetStreamSettingsResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

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
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

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
