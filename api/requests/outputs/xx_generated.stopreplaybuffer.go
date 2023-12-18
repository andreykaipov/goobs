// This file has been automatically generated. Don't edit it.

package outputs

import api "github.com/andreykaipov/goobs/api"

// Represents the request body for the StopReplayBuffer request.
type StopReplayBufferParams struct{}

// Returns the associated request.
func (o *StopReplayBufferParams) GetRequestName() string {
	return "StopReplayBuffer"
}

// Represents the response body for the StopReplayBuffer request.
type StopReplayBufferResponse struct {
	api.ResponseCommon
}

// Stops the replay buffer output.
func (c *Client) StopReplayBuffer(paramss ...*StopReplayBufferParams) (*StopReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopReplayBufferParams{{}}
	}
	params := paramss[0]
	data := &StopReplayBufferResponse{}
	return data, c.SendRequest(params, data)
}
