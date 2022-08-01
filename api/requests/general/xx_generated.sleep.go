// This file has been automatically generated. Don't edit it.

package general

/*
SleepParams represents the params body for the "Sleep" request.
Sleeps for a time duration or number of frames. Only available in request batches with types `SERIAL_REALTIME` or `SERIAL_FRAME`.
*/
type SleepParams struct {
	// Number of frames to sleep for (if `SERIAL_FRAME` mode)
	SleepFrames float64 `json:"sleepFrames,omitempty"`

	// Number of milliseconds to sleep for (if `SERIAL_REALTIME` mode)
	SleepMillis float64 `json:"sleepMillis,omitempty"`
}

/*
SleepResponse represents the response body for the "Sleep" request.
Sleeps for a time duration or number of frames. Only available in request batches with types `SERIAL_REALTIME` or `SERIAL_FRAME`.
*/
type SleepResponse struct{}

// Sleep sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) Sleep(params *SleepParams) (*SleepResponse, error) {
	resp, err := c.SendRequest("Sleep", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SleepResponse), nil
}
