// This file has been automatically generated. Don't edit it.

package replaybuffer

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartStopReplayBufferParams represents the params body for the "StartStopReplayBuffer" request.
Toggle the Replay Buffer on/off (depending on the current state of the replay buffer).
Since 4.2.0.
*/
type StartStopReplayBufferParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StartStopReplayBuffer".
func (o *StartStopReplayBufferParams) GetSelfName() string {
	return "StartStopReplayBuffer"
}

/*
StartStopReplayBufferResponse represents the response body for the "StartStopReplayBuffer" request.
Toggle the Replay Buffer on/off (depending on the current state of the replay buffer).
Since v4.2.0.
*/
type StartStopReplayBufferResponse struct {
	requests.ResponseBasic
}

// StartStopReplayBuffer sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) StartStopReplayBuffer(
	paramss ...*StartStopReplayBufferParams,
) (*StartStopReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartStopReplayBufferParams{{}}
	}
	params := paramss[0]
	data := &StartStopReplayBufferResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
