// This file has been automatically generated. Don't edit it.

package general

import api "github.com/andreykaipov/goobs/api"

/*
GetVersionParams represents the params body for the "GetVersion" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVersion.
*/
type GetVersionParams struct {
	api.Params
}

func (o *GetVersionParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetVersionParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetVersionResponse represents the response body for the "GetVersion" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVersion.
*/
type GetVersionResponse struct {
	api.Response

	// List of available request types, formatted as a comma-separated list string (e.g. :
	// "Method1,Method2,Method3").
	AvailableRequests string `json:"available-requests"`

	// OBS Studio program version.
	ObsStudioVersion string `json:"obs-studio-version"`

	// obs-websocket plugin version.
	ObsWebsocketVersion string `json:"obs-websocket-version"`

	// OBSRemote compatible API version. Fixed to 1.1 for retrocompatibility.
	Version float64 `json:"version"`
}

func (o *GetVersionResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetVersionResponse) GetStatus() string {
	return o.Status
}
func (o *GetVersionResponse) GetError() string {
	return o.Error
}

/*
GetAuthRequiredParams represents the params body for the "GetAuthRequired" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetAuthRequired.
*/
type GetAuthRequiredParams struct {
	api.Params
}

func (o *GetAuthRequiredParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetAuthRequiredParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetAuthRequiredResponse represents the response body for the "GetAuthRequired" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetAuthRequired.
*/
type GetAuthRequiredResponse struct {
	api.Response

	// Indicates whether authentication is required.
	AuthRequired bool `json:"authRequired"`

	Challenge string `json:"challenge"`

	Salt string `json:"salt"`
}

func (o *GetAuthRequiredResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetAuthRequiredResponse) GetStatus() string {
	return o.Status
}
func (o *GetAuthRequiredResponse) GetError() string {
	return o.Error
}

/*
AuthenticateParams represents the params body for the "Authenticate" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#Authenticate.
*/
type AuthenticateParams struct {
	api.Params

	// Response to the auth challenge (see "Authentication" for more information).
	Auth string `json:"auth"`
}

func (o *AuthenticateParams) GetRequestType() string {
	return o.RequestType
}
func (o *AuthenticateParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
AuthenticateResponse represents the response body for the "Authenticate" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#Authenticate.
*/
type AuthenticateResponse struct {
	api.Response
}

func (o *AuthenticateResponse) GetMessageID() string {
	return o.MessageID
}
func (o *AuthenticateResponse) GetStatus() string {
	return o.Status
}
func (o *AuthenticateResponse) GetError() string {
	return o.Error
}

/*
SetHeartbeatParams represents the params body for the "SetHeartbeat" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetHeartbeat.
*/
type SetHeartbeatParams struct {
	api.Params

	// Starts/Stops emitting heartbeat messages
	Enable bool `json:"enable"`
}

func (o *SetHeartbeatParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetHeartbeatParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetHeartbeatResponse represents the response body for the "SetHeartbeat" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetHeartbeat.
*/
type SetHeartbeatResponse struct {
	api.Response
}

func (o *SetHeartbeatResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetHeartbeatResponse) GetStatus() string {
	return o.Status
}
func (o *SetHeartbeatResponse) GetError() string {
	return o.Error
}

/*
SetFilenameFormattingParams represents the params body for the "SetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetFilenameFormatting.
*/
type SetFilenameFormattingParams struct {
	api.Params

	// Filename formatting string to set.
	FilenameFormatting string `json:"filename-formatting"`
}

func (o *SetFilenameFormattingParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetFilenameFormattingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetFilenameFormattingResponse represents the response body for the "SetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetFilenameFormatting.
*/
type SetFilenameFormattingResponse struct {
	api.Response
}

func (o *SetFilenameFormattingResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetFilenameFormattingResponse) GetStatus() string {
	return o.Status
}
func (o *SetFilenameFormattingResponse) GetError() string {
	return o.Error
}

/*
GetFilenameFormattingParams represents the params body for the "GetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetFilenameFormatting.
*/
type GetFilenameFormattingParams struct {
	api.Params
}

func (o *GetFilenameFormattingParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetFilenameFormattingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetFilenameFormattingResponse represents the response body for the "GetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetFilenameFormatting.
*/
type GetFilenameFormattingResponse struct {
	api.Response

	// Current filename formatting string.
	FilenameFormatting string `json:"filename-formatting"`
}

func (o *GetFilenameFormattingResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetFilenameFormattingResponse) GetStatus() string {
	return o.Status
}
func (o *GetFilenameFormattingResponse) GetError() string {
	return o.Error
}
