// This file has been automatically generated. Don't edit it.

package general

/*
GetVersionParams represents the params body for the "GetVersion" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVersion.
*/
type GetVersionParams struct{}

/*
GetVersionResponse represents the response body for the "GetVersion" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVersion.
*/
type GetVersionResponse struct {
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

/*
GetAuthRequiredParams represents the params body for the "GetAuthRequired" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetAuthRequired.
*/
type GetAuthRequiredParams struct{}

/*
GetAuthRequiredResponse represents the response body for the "GetAuthRequired" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetAuthRequired.
*/
type GetAuthRequiredResponse struct {
	// Indicates whether authentication is required.
	AuthRequired bool `json:"authRequired"`

	Challenge string `json:"challenge"`

	Salt string `json:"salt"`
}

/*
AuthenticateParams represents the params body for the "Authenticate" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#Authenticate.
*/
type AuthenticateParams struct {
	// Response to the auth challenge (see "Authentication" for more information).
	Auth string `json:"auth"`
}

/*
AuthenticateResponse represents the response body for the "Authenticate" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#Authenticate.
*/
type AuthenticateResponse struct{}

/*
SetHeartbeatParams represents the params body for the "SetHeartbeat" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetHeartbeat.
*/
type SetHeartbeatParams struct {
	// Starts/Stops emitting heartbeat messages
	Enable bool `json:"enable"`
}

/*
SetHeartbeatResponse represents the response body for the "SetHeartbeat" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetHeartbeat.
*/
type SetHeartbeatResponse struct{}

/*
SetFilenameFormattingParams represents the params body for the "SetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetFilenameFormatting.
*/
type SetFilenameFormattingParams struct {
	// Filename formatting string to set.
	FilenameFormatting string `json:"filename-formatting"`
}

/*
SetFilenameFormattingResponse represents the response body for the "SetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetFilenameFormatting.
*/
type SetFilenameFormattingResponse struct{}

/*
GetFilenameFormattingParams represents the params body for the "GetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetFilenameFormatting.
*/
type GetFilenameFormattingParams struct{}

/*
GetFilenameFormattingResponse represents the response body for the "GetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetFilenameFormatting.
*/
type GetFilenameFormattingResponse struct {
	// Current filename formatting string.
	FilenameFormatting string `json:"filename-formatting"`
}
