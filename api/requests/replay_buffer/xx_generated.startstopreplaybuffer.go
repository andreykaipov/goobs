// This file has been automatically generated. Don't edit it.

package replaybuffer

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartStopReplayBufferParams represents the params body for the "StartStopReplayBuffer" request.
Toggle the Replay Buffer on/off.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopReplayBuffer.
*/
type StartStopReplayBufferParams struct {
	requests.ParamsBasic
}

// Name just returns "StartStopReplayBuffer".
func (o *StartStopReplayBufferParams) Name() string {
	return "StartStopReplayBuffer"
}

/*
StartStopReplayBufferResponse represents the response body for the "StartStopReplayBuffer" request.
Toggle the Replay Buffer on/off.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopReplayBuffer.
*/
type StartStopReplayBufferResponse struct {
	requests.ResponseBasic
}

// StartStopReplayBuffer sends the corresponding request to the connected OBS WebSockets server.
// Note the variadic arguments as this request doesn't require any parameters.
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
