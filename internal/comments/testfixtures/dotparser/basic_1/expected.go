package testfixtures

type StartStreamingRequest struct {
	Stream *Stream `json:"stream,omitempty"`
}

type Settings struct {
	Key string `json:"key"`

	Password string `json:"password"`

	Server string `json:"server"`

	UseAuth bool `json:"use-auth"`

	Username string `json:"username"`
}

type Stream struct {
	Metadata map[string]interface{} `json:"metadata"`

	Settings *Settings `json:"settings,omitempty"`

	Type int `json:"type"`
}
