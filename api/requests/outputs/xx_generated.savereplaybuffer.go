// This file has been automatically generated. Don't edit it.

package outputs

/*
SaveReplayBufferParams represents the params body for the "SaveReplayBuffer" request.
Saves the contents of the replay buffer output.
*/
type SaveReplayBufferParams struct{}

/*
SaveReplayBufferResponse represents the response body for the "SaveReplayBuffer" request.
Saves the contents of the replay buffer output.
*/
type SaveReplayBufferResponse struct{}

// SaveReplayBuffer sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) SaveReplayBuffer(paramss ...*SaveReplayBufferParams) (*SaveReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*SaveReplayBufferParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("SaveReplayBuffer", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SaveReplayBufferResponse), nil
}
