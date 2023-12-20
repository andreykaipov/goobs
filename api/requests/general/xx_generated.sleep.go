// This file has been automatically generated. Don't edit it.

package general

// Represents the request body for the Sleep request.
type SleepParams struct {
	// Number of frames to sleep for (if `SERIAL_FRAME` mode)
	SleepFrames *float64 `json:"sleepFrames,omitempty"`

	// Number of milliseconds to sleep for (if `SERIAL_REALTIME` mode)
	SleepMillis *float64 `json:"sleepMillis,omitempty"`
}

func NewSleepParams() *SleepParams {
	return &SleepParams{}
}
func (o *SleepParams) WithSleepFrames(x float64) *SleepParams {
	o.SleepFrames = &x
	return o
}
func (o *SleepParams) WithSleepMillis(x float64) *SleepParams {
	o.SleepMillis = &x
	return o
}

// Returns the associated request.
func (o *SleepParams) GetRequestName() string {
	return "Sleep"
}

// Represents the response body for the Sleep request.
type SleepResponse struct {
	_response
}

// Sleeps for a time duration or number of frames. Only available in request batches with types `SERIAL_REALTIME` or
// `SERIAL_FRAME`.
func (c *Client) Sleep(paramss ...*SleepParams) (*SleepResponse, error) {
	if len(paramss) == 0 {
		paramss = []*SleepParams{{}}
	}
	params := paramss[0]
	data := &SleepResponse{}
	return data, c.client.SendRequest(params, data)
}
