// This file has been automatically generated. Don't edit it.

package streaming

type GetStreamingStatusParams struct{}
type GetStreamingStatusResponse struct {
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
type StartStopStreamingParams struct{}
type StartStopStreamingResponse struct{}
type StartStreamingParams struct {
	Stream struct {
		// Adds the given object parameters as encoded query string parameters to the 'key' of the RTMP stream. Used to pass data to the RTMP service about the streaming. May be any String, Numeric, or Boolean field.
		Metadata map[string]interface{} `json:"metadata"`
		Settings struct {
			// The publish key of the stream.
			Key string `json:"key"`
			// If authentication is enabled, the password for the streaming server. Ignored if `use-auth` is not set to `true`.
			Password string `json:"password"`
			// The publish URL.
			Server string `json:"server"`
			// Indicates whether authentication should be used when connecting to the streaming server.
			UseAuth bool `json:"use-auth"`
			// If authentication is enabled, the username for the streaming server. Ignored if `use-auth` is not set to `true`.
			Username string `json:"username"`
		} `json:"settings"`
		// If specified ensures the type of stream matches the given type (usually 'rtmp_custom' or 'rtmp_common'). If the currently configured stream type does not match the given stream type, all settings must be specified in the `settings` object or an error will occur when starting the stream.
		Type string `json:"type"`
	} `json:"stream"`
}
type StartStreamingResponse struct{}
type StopStreamingParams struct{}
type StopStreamingResponse struct{}
type SetStreamSettingsParams struct {
	// Persist the settings to disk.
	Save     bool `json:"save"`
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
type SetStreamSettingsResponse struct{}
type GetStreamSettingsParams struct{}
type GetStreamSettingsResponse struct {
	Settings struct {
		// The publish key of the stream.
		Key string `json:"key"`
		// The password to use when accessing the streaming server. Only present if `use-auth` is `true`.
		Password string `json:"password"`
		// The publish URL.
		Server string `json:"server"`
		// Indicates whether authentication should be used when connecting to the streaming server.
		UseAuth bool `json:"use-auth"`
		// The username to use when accessing the streaming server. Only present if `use-auth` is `true`.
		Username string `json:"username"`
	} `json:"settings"`
	// The type of streaming service configuration. Possible values: 'rtmp_custom' or 'rtmp_common'.
	Type string `json:"type"`
}
type SaveStreamSettingsParams struct{}
type SaveStreamSettingsResponse struct{}
