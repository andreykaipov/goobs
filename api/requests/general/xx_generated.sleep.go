// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SleepParams represents the params body for the "Sleep" request.
Sleeps for a time duration or number of frames. Only available in request batches with types `SERIAL_REALTIME` or `SERIAL_FRAME`.
*/
type SleepParams struct {
	requests.ParamsBasic

	// Number of frames to sleep for (if `SERIAL_FRAME` mode)
	SleepFrames float64 `json:"sleepFrames,omitempty"`

	// Number of milliseconds to sleep for (if `SERIAL_REALTIME` mode)
	SleepMillis float64 `json:"sleepMillis,omitempty"`
}

// GetSelfName just returns "Sleep".
func (o *SleepParams) GetSelfName() string {
	return "Sleep"
}

/*
SleepResponse represents the response body for the "Sleep" request.
Sleeps for a time duration or number of frames. Only available in request batches with types `SERIAL_REALTIME` or `SERIAL_FRAME`.
*/
type SleepResponse struct {
	requests.ResponseBasic
}

// Sleep sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) Sleep(params *SleepParams) (*SleepResponse, error) {
	data := &SleepResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
