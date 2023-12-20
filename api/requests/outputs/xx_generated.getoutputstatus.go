// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the GetOutputStatus request.
type GetOutputStatusParams struct {
	// Output name
	OutputName *string `json:"outputName,omitempty"`
}

func NewGetOutputStatusParams() *GetOutputStatusParams {
	return &GetOutputStatusParams{}
}
func (o *GetOutputStatusParams) WithOutputName(x string) *GetOutputStatusParams {
	o.OutputName = &x
	return o
}

// Returns the associated request.
func (o *GetOutputStatusParams) GetRequestName() string {
	return "GetOutputStatus"
}

// Represents the response body for the GetOutputStatus request.
type GetOutputStatusResponse struct {
	_response

	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`

	// Number of bytes sent by the output
	OutputBytes float64 `json:"outputBytes,omitempty"`

	// Congestion of the output
	OutputCongestion float64 `json:"outputCongestion,omitempty"`

	// Current duration in milliseconds for the output
	OutputDuration float64 `json:"outputDuration,omitempty"`

	// Whether the output is reconnecting
	OutputReconnecting bool `json:"outputReconnecting,omitempty"`

	// Number of frames skipped by the output's process
	OutputSkippedFrames float64 `json:"outputSkippedFrames,omitempty"`

	// Current formatted timecode string for the output
	OutputTimecode string `json:"outputTimecode,omitempty"`

	// Total number of frames delivered by the output's process
	OutputTotalFrames float64 `json:"outputTotalFrames,omitempty"`
}

// Gets the status of an output.
func (c *Client) GetOutputStatus(params *GetOutputStatusParams) (*GetOutputStatusResponse, error) {
	data := &GetOutputStatusResponse{}
	return data, c.client.SendRequest(params, data)
}
