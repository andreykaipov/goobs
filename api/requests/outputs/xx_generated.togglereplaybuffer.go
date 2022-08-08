// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the ToggleReplayBuffer request.
type ToggleReplayBufferParams struct{}

// Returns the associated request.
func (o *ToggleReplayBufferParams) GetRequestName() string {
	return "ToggleReplayBuffer"
}

// Represents the response body for the ToggleReplayBuffer request.
type ToggleReplayBufferResponse struct {
	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`
}

// Toggles the state of the replay buffer output.
func (c *Client) ToggleReplayBuffer(paramss ...*ToggleReplayBufferParams) (*ToggleReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleReplayBufferParams{{}}
	}
	params := paramss[0]
	data := &ToggleReplayBufferResponse{}
	return data, c.SendRequest(params, data)
}
