// This file has been automatically generated. Don't edit it.

package outputs

/*
StartReplayBufferParams represents the params body for the "StartReplayBuffer" request.
Starts the replay buffer output.
*/
type StartReplayBufferParams struct{}

/*
StartReplayBufferResponse represents the response body for the "StartReplayBuffer" request.
Starts the replay buffer output.
*/
type StartReplayBufferResponse struct{}

// StartReplayBuffer sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) StartReplayBuffer(paramss ...*StartReplayBufferParams) (*StartReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartReplayBufferParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("StartReplayBuffer", params)
	if err != nil {
		return nil, err
	}
	return resp.(*StartReplayBufferResponse), nil
}
