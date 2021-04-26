// This file has been automatically generated. Don't edit it.

package general

type GetVersionParams struct{}
type GetVersionResponse struct {
	// List of available request types, formatted as a comma-separated list string (e.g. : "Method1,Method2,Method3").
	AvailableRequests string `json:"available-requests"`
	// OBS Studio program version.
	ObsStudioVersion string `json:"obs-studio-version"`
	// obs-websocket plugin version.
	ObsWebsocketVersion string `json:"obs-websocket-version"`
	// OBSRemote compatible API version. Fixed to 1.1 for retrocompatibility.
	Version float64 `json:"version"`
}
type GetAuthRequiredParams struct{}
type GetAuthRequiredResponse struct {
	// Indicates whether authentication is required.
	AuthRequired bool   `json:"authRequired"`
	Challenge    string `json:"challenge"`
	Salt         string `json:"salt"`
}
type AuthenticateParams struct {
	// Response to the auth challenge (see "Authentication" for more information).
	Auth string `json:"auth"`
}
type AuthenticateResponse struct{}
type SetHeartbeatParams struct {
	// Starts/Stops emitting heartbeat messages
	Enable bool `json:"enable"`
}
type SetHeartbeatResponse struct{}
type SetFilenameFormattingParams struct {
	// Filename formatting string to set.
	FilenameFormatting string `json:"filename-formatting"`
}
type SetFilenameFormattingResponse struct{}
type GetFilenameFormattingParams struct{}
type GetFilenameFormattingResponse struct {
	// Current filename formatting string.
	FilenameFormatting string `json:"filename-formatting"`
}
