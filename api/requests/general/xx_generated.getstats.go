// This file has been automatically generated. Don't edit it.

package general

// Represents the request body for the GetStats request.
type GetStatsParams struct{}

// Returns the associated request.
func (o *GetStatsParams) GetRequestName() string {
	return "GetStats"
}

// Represents the response body for the GetStats request.
type GetStatsResponse struct {
	// Current FPS being rendered
	ActiveFps float64 `json:"activeFps,omitempty"`

	// Available disk space on the device being used for recording storage
	AvailableDiskSpace float64 `json:"availableDiskSpace,omitempty"`

	// Average time in milliseconds that OBS is taking to render a frame
	AverageFrameRenderTime float64 `json:"averageFrameRenderTime,omitempty"`

	// Current CPU usage in percent
	CpuUsage float64 `json:"cpuUsage,omitempty"`

	// Amount of memory in MB currently being used by OBS
	MemoryUsage float64 `json:"memoryUsage,omitempty"`

	// Number of frames skipped by OBS in the output thread
	OutputSkippedFrames float64 `json:"outputSkippedFrames,omitempty"`

	// Total number of frames outputted by the output thread
	OutputTotalFrames float64 `json:"outputTotalFrames,omitempty"`

	// Number of frames skipped by OBS in the render thread
	RenderSkippedFrames float64 `json:"renderSkippedFrames,omitempty"`

	// Total number of frames outputted by the render thread
	RenderTotalFrames float64 `json:"renderTotalFrames,omitempty"`

	// Total number of messages received by obs-websocket from the client
	WebSocketSessionIncomingMessages float64 `json:"webSocketSessionIncomingMessages,omitempty"`

	// Total number of messages sent by obs-websocket to the client
	WebSocketSessionOutgoingMessages float64 `json:"webSocketSessionOutgoingMessages,omitempty"`
}

// Gets statistics about OBS, obs-websocket, and the current session.
func (c *Client) GetStats(paramss ...*GetStatsParams) (*GetStatsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStatsParams{{}}
	}
	params := paramss[0]
	data := &GetStatsResponse{}
	return data, c.SendRequest(params, data)
}
