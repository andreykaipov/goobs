package typedefs

// StreamSettings specifies the stream settings.
type StreamSettings struct {
	// The publish key of the stream.
	Key string `json:"key,omitempty"`

	// The password to use when accessing the streaming server. Only
	// relevant if `use_auth` is `true`.
	Password string `json:"password,omitempty"`

	// The publish URL.
	Server string `json:"server,omitempty"`

	// Indicates whether authentication should be used when connecting to
	// the streaming server.
	UseAuth bool `json:"use_auth"`

	// The username to use when accessing the streaming server. Only
	// relevant if `use_auth` is `true`.
	Username string `json:"username,omitempty"`
}
