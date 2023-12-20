// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the StartReplayBuffer request.
type StartReplayBufferParams struct{}

// Returns the associated request.
func (o *StartReplayBufferParams) GetRequestName() string {
	return "StartReplayBuffer"
}

// Represents the response body for the StartReplayBuffer request.
type StartReplayBufferResponse struct {
	_response
}

// Starts the replay buffer output.
func (c *Client) StartReplayBuffer(paramss ...*StartReplayBufferParams) (*StartReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartReplayBufferParams{{}}
	}
	params := paramss[0]
	data := &StartReplayBufferResponse{}
	return data, c.client.SendRequest(params, data)
}
