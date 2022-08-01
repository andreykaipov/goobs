// This file has been automatically generated. Don't edit it.

package outputs

/*
StopReplayBufferParams represents the params body for the "StopReplayBuffer" request.
Stops the replay buffer output.
*/
type StopReplayBufferParams struct{}

/*
StopReplayBufferResponse represents the response body for the "StopReplayBuffer" request.
Stops the replay buffer output.
*/
type StopReplayBufferResponse struct{}

// StopReplayBuffer sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) StopReplayBuffer(paramss ...*StopReplayBufferParams) (*StopReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopReplayBufferParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("StopReplayBuffer", params)
	if err != nil {
		return nil, err
	}
	return resp.(*StopReplayBufferResponse), nil
}
