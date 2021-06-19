package typedefs

// Settings specifies the stream settings.
type Settings struct {
	// The publish key of the stream.
	Key string `json:"key"`

	// The password to use when accessing the streaming server. Only
	// relevant if `use_auth` is `true`.
	Password string `json:"password"`

	// The publish URL.
	Server string `json:"server"`

	// Indicates whether authentication should be used when connecting to
	// the streaming server.
	UseAuth bool `json:"use_auth"`

	// The username to use when accessing the streaming server. Only
	// relevant if `use_auth` is `true`.
	Username string `json:"username"`
}
