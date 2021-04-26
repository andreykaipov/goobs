// This file has been automatically generated. Don't edit it.

package streaming

type GetStreamingStatusParams struct{}
type GetStreamingStatusResponse struct {
	PreviewOnly    bool   `json:"preview-only"`
	RecTimecode    string `json:"rec-timecode"`
	Recording      bool   `json:"recording"`
	StreamTimecode string `json:"stream-timecode"`
	Streaming      bool   `json:"streaming"`
}
type StartStopStreamingParams struct{}
type StartStopStreamingResponse struct{}
type StartStreamingParams struct {
	Stream struct {
		Metadata map[string]interface{} `json:"metadata"`
		Settings struct {
			Key      string `json:"key"`
			Password string `json:"password"`
			Server   string `json:"server"`
			UseAuth  bool   `json:"use-auth"`
			Username string `json:"username"`
		} `json:"settings"`
		Type string `json:"type"`
	} `json:"stream"`
}
type StartStreamingResponse struct{}
type StopStreamingParams struct{}
type StopStreamingResponse struct{}
type SetStreamSettingsParams struct {
	Save     bool `json:"save"`
	Settings struct {
		Key      string `json:"key"`
		Password string `json:"password"`
		Server   string `json:"server"`
		UseAuth  bool   `json:"use-auth"`
		Username string `json:"username"`
	} `json:"settings"`
	Type string `json:"type"`
}
type SetStreamSettingsResponse struct{}
type GetStreamSettingsParams struct{}
type GetStreamSettingsResponse struct {
	Settings struct {
		Key      string `json:"key"`
		Password string `json:"password"`
		Server   string `json:"server"`
		UseAuth  bool   `json:"use-auth"`
		Username string `json:"username"`
	} `json:"settings"`
	Type string `json:"type"`
}
type SaveStreamSettingsParams struct{}
type SaveStreamSettingsResponse struct{}
