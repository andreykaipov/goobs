// This file has been automatically generated. Don't edit it.

package general

type GetVersionParams struct{}
type GetVersionResponse struct {
	AvailableRequests   string  `json:"available-requests"`
	ObsStudioVersion    string  `json:"obs-studio-version"`
	ObsWebsocketVersion string  `json:"obs-websocket-version"`
	Version             float64 `json:"version"`
}
type GetAuthRequiredParams struct{}
type GetAuthRequiredResponse struct {
	AuthRequired bool   `json:"authRequired"`
	Challenge    string `json:"challenge"`
	Salt         string `json:"salt"`
}
type AuthenticateParams struct {
	Auth string `json:"auth"`
}
type AuthenticateResponse struct{}
type SetHeartbeatParams struct {
	Enable bool `json:"enable"`
}
type SetHeartbeatResponse struct{}
type SetFilenameFormattingParams struct {
	FilenameFormatting string `json:"filename-formatting"`
}
type SetFilenameFormattingResponse struct{}
type GetFilenameFormattingParams struct{}
type GetFilenameFormattingResponse struct {
	FilenameFormatting string `json:"filename-formatting"`
}
