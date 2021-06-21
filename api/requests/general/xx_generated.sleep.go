// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SleepParams represents the params body for the "Sleep" request.
Waits for the specified duration. Designed to be used in `ExecuteBatch` operations.
Since 4.9.1.
*/
type SleepParams struct {
	requests.ParamsBasic

	// Delay in milliseconds to wait before continuing.
	SleepMillis int `json:"sleepMillis,omitempty"`
}

// GetSelfName just returns "Sleep".
func (o *SleepParams) GetSelfName() string {
	return "Sleep"
}

/*
SleepResponse represents the response body for the "Sleep" request.
Waits for the specified duration. Designed to be used in `ExecuteBatch` operations.
Since v4.9.1.
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
