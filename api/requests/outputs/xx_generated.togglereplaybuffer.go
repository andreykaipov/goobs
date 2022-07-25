// This file has been automatically generated. Don't edit it.

package outputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ToggleReplayBufferParams represents the params body for the "ToggleReplayBuffer" request.
Toggles the state of the replay buffer output.
*/
type ToggleReplayBufferParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "ToggleReplayBuffer".
func (o *ToggleReplayBufferParams) GetSelfName() string {
	return "ToggleReplayBuffer"
}

/*
ToggleReplayBufferResponse represents the response body for the "ToggleReplayBuffer" request.
Toggles the state of the replay buffer output.
*/
type ToggleReplayBufferResponse struct {
	requests.ResponseBasic

	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`
}

// ToggleReplayBuffer sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) ToggleReplayBuffer(paramss ...*ToggleReplayBufferParams) (*ToggleReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleReplayBufferParams{{}}
	}
	params := paramss[0]
	data := &ToggleReplayBufferResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
