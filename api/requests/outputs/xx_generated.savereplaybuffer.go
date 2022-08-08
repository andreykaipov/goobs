// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the SaveReplayBuffer request.
type SaveReplayBufferParams struct{}

// Returns the associated request.
func (o *SaveReplayBufferParams) GetRequestName() string {
	return "SaveReplayBuffer"
}

// Represents the response body for the SaveReplayBuffer request.
type SaveReplayBufferResponse struct{}

// Saves the contents of the replay buffer output.
func (c *Client) SaveReplayBuffer(paramss ...*SaveReplayBufferParams) (*SaveReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*SaveReplayBufferParams{{}}
	}
	params := paramss[0]
	data := &SaveReplayBufferResponse{}
	return data, c.SendRequest(params, data)
}
