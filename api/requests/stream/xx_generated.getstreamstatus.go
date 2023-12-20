// This file has been automatically generated. Don't edit it.

package stream

// Represents the request body for the GetStreamStatus request.
type GetStreamStatusParams struct{}

// Returns the associated request.
func (o *GetStreamStatusParams) GetRequestName() string {
	return "GetStreamStatus"
}

// Represents the response body for the GetStreamStatus request.
type GetStreamStatusResponse struct {
	_response

	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`

	// Number of bytes sent by the output
	OutputBytes float64 `json:"outputBytes,omitempty"`

	// Congestion of the output
	OutputCongestion float64 `json:"outputCongestion,omitempty"`

	// Current duration in milliseconds for the output
	OutputDuration float64 `json:"outputDuration,omitempty"`

	// Whether the output is currently reconnecting
	OutputReconnecting bool `json:"outputReconnecting,omitempty"`

	// Number of frames skipped by the output's process
	OutputSkippedFrames float64 `json:"outputSkippedFrames,omitempty"`

	// Current formatted timecode string for the output
	OutputTimecode string `json:"outputTimecode,omitempty"`

	// Total number of frames delivered by the output's process
	OutputTotalFrames float64 `json:"outputTotalFrames,omitempty"`
}

// Gets the status of the stream output.
func (c *Client) GetStreamStatus(paramss ...*GetStreamStatusParams) (*GetStreamStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStreamStatusParams{{}}
	}
	params := paramss[0]
	data := &GetStreamStatusResponse{}
	return data, c.client.SendRequest(params, data)
}
