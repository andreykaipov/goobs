// This file has been automatically generated. Don't edit it.

package general

// GetVersionParams contains the request body for the
// [GetVersion](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVersion)
// request.
type GetVersionParams struct{}

// GetVersionResponse contains the request body for the
// [GetVersion](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVersion)
// request.
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

// GetAuthRequiredParams contains the request body for the
// [GetAuthRequired](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetAuthRequired)
// request.
type GetAuthRequiredParams struct{}

// GetAuthRequiredResponse contains the request body for the
// [GetAuthRequired](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetAuthRequired)
// request.
type GetAuthRequiredResponse struct {
	// Indicates whether authentication is required.
	AuthRequired bool `json:"authRequired"`

	Challenge string `json:"challenge"`

	Salt string `json:"salt"`
}

// AuthenticateParams contains the request body for the
// [Authenticate](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#Authenticate)
// request.
type AuthenticateParams struct {
	// Response to the auth challenge (see "Authentication" for more information).
	Auth string `json:"auth"`
}

// AuthenticateResponse contains the request body for the
// [Authenticate](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#Authenticate)
// request.
type AuthenticateResponse struct{}

// SetHeartbeatParams contains the request body for the
// [SetHeartbeat](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetHeartbeat)
// request.
type SetHeartbeatParams struct {
	// Starts/Stops emitting heartbeat messages
	Enable bool `json:"enable"`
}

// SetHeartbeatResponse contains the request body for the
// [SetHeartbeat](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetHeartbeat)
// request.
type SetHeartbeatResponse struct{}

// SetFilenameFormattingParams contains the request body for the
// [SetFilenameFormatting](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetFilenameFormatting)
// request.
type SetFilenameFormattingParams struct {
	// Filename formatting string to set.
	FilenameFormatting string `json:"filename-formatting"`
}

// SetFilenameFormattingResponse contains the request body for the
// [SetFilenameFormatting](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetFilenameFormatting)
// request.
type SetFilenameFormattingResponse struct{}

// GetFilenameFormattingParams contains the request body for the
// [GetFilenameFormatting](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetFilenameFormatting)
// request.
type GetFilenameFormattingParams struct{}

// GetFilenameFormattingResponse contains the request body for the
// [GetFilenameFormatting](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetFilenameFormatting)
// request.
type GetFilenameFormattingResponse struct {
	// Current filename formatting string.
	FilenameFormatting string `json:"filename-formatting"`
}
