// This file has been automatically generated. Don't edit it.

package general

// Represents the request body for the GetVersion request.
type GetVersionParams struct{}

// Returns the associated request.
func (o *GetVersionParams) GetRequestName() string {
	return "GetVersion"
}

// Represents the response body for the GetVersion request.
type GetVersionResponse struct {
	// Array of available RPC requests for the currently negotiated RPC version
	AvailableRequests []string `json:"availableRequests,omitempty"`

	// Current OBS Studio version
	ObsVersion string `json:"obsVersion,omitempty"`

	// Current obs-websocket version
	ObsWebSocketVersion string `json:"obsWebSocketVersion,omitempty"`

	// Name of the platform. Usually `windows`, `macos`, or `ubuntu` (linux flavor). Not guaranteed to be any of those
	Platform string `json:"platform,omitempty"`

	// Description of the platform, like `Windows 10 (10.0)`
	PlatformDescription string `json:"platformDescription,omitempty"`

	// Current latest obs-websocket RPC version
	RpcVersion float64 `json:"rpcVersion,omitempty"`

	// Image formats available in `GetSourceScreenshot` and `SaveSourceScreenshot` requests.
	SupportedImageFormats []string `json:"supportedImageFormats,omitempty"`
}

// Gets data about the current plugin and RPC version.
func (c *Client) GetVersion(paramss ...*GetVersionParams) (*GetVersionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetVersionParams{{}}
	}
	params := paramss[0]
	data := &GetVersionResponse{}
	return data, c.SendRequest(params, data)
}
