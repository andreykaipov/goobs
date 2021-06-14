// This file has been automatically generated. Don't edit it.

package replaybuffer

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StopReplayBufferParams represents the params body for the "StopReplayBuffer" request.
Stop recording into the Replay Buffer.
Will return an `error` if the Replay Buffer is not active.
Since 4.2.0.
*/
type StopReplayBufferParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StopReplayBuffer".
func (o *StopReplayBufferParams) GetSelfName() string {
	return "StopReplayBuffer"
}

/*
StopReplayBufferResponse represents the response body for the "StopReplayBuffer" request.
Stop recording into the Replay Buffer.
Will return an `error` if the Replay Buffer is not active.
Since v4.2.0.
*/
type StopReplayBufferResponse struct {
	requests.ResponseBasic
}

// StopReplayBuffer sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) StopReplayBuffer(paramss ...*StopReplayBufferParams) (*StopReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopReplayBufferParams{{}}
	}
	params := paramss[0]
	data := &StopReplayBufferResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
