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

func (o *GetStreamingStatusParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *GetStreamingStatusResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetStreamingStatusResponse) GetStatus() string {
	return o.Status
}
func (o *GetStreamingStatusResponse) GetError() string {
	return o.Error
}

/*
StartStopStreamingParams represents the params body for the "StartStopStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopStreaming.
*/
type StartStopStreamingParams struct {
	requests.Params
}

func (o *StartStopStreamingParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *StartStopStreamingResponse) GetMessageID() string {
	return o.MessageID
}
func (o *StartStopStreamingResponse) GetStatus() string {
	return o.Status
}
func (o *StartStopStreamingResponse) GetError() string {
	return o.Error
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

func (o *StartStreamingParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *StartStreamingResponse) GetMessageID() string {
	return o.MessageID
}
func (o *StartStreamingResponse) GetStatus() string {
	return o.Status
}
func (o *StartStreamingResponse) GetError() string {
	return o.Error
}

/*
StopStreamingParams represents the params body for the "StopStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopStreaming.
*/
type StopStreamingParams struct {
	requests.Params
}

func (o *StopStreamingParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *StopStreamingResponse) GetMessageID() string {
	return o.MessageID
}
func (o *StopStreamingResponse) GetStatus() string {
	return o.Status
}
func (o *StopStreamingResponse) GetError() string {
	return o.Error
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

func (o *SetStreamSettingsParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *SetStreamSettingsResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetStreamSettingsResponse) GetStatus() string {
	return o.Status
}
func (o *SetStreamSettingsResponse) GetError() string {
	return o.Error
}

/*
GetStreamSettingsParams represents the params body for the "GetStreamSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStreamSettings.
*/
type GetStreamSettingsParams struct {
	requests.Params
}

func (o *GetStreamSettingsParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *GetStreamSettingsResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetStreamSettingsResponse) GetStatus() string {
	return o.Status
}
func (o *GetStreamSettingsResponse) GetError() string {
	return o.Error
}

/*
SaveStreamSettingsParams represents the params body for the "SaveStreamSettings" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SaveStreamSettings.
*/
type SaveStreamSettingsParams struct {
	requests.Params
}

func (o *SaveStreamSettingsParams) GetRequestType() string {
	return o.RequestType
}
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

func (o *SaveStreamSettingsResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SaveStreamSettingsResponse) GetStatus() string {
	return o.Status
}
func (o *SaveStreamSettingsResponse) GetError() string {
	return o.Error
}
