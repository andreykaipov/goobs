package testfixtures

type StartStreamingRequest struct {
	Stream struct {
		Metadata map[string]interface{} `json:"metadata"`

		Settings struct {
			Key string `json:"key"`

			Password string `json:"password"`

			Server string `json:"server"`

			UseAuth bool `json:"use-auth"`

			Username string `json:"username"`
		} `json:"settings"`

		Type int `json:"type"`
	} `json:"stream"`
}
