// This file has been automatically generated. Don't edit it.

package replaybuffer

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartReplayBufferParams represents the params body for the "StartReplayBuffer" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartReplayBuffer.
*/
type StartReplayBufferParams struct {
	requests.ParamsBasic
}

// Name just returns "StartReplayBuffer".
func (o *StartReplayBufferParams) Name() string {
	return "StartReplayBuffer"
}

/*
StartReplayBufferResponse represents the response body for the "StartReplayBuffer" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartReplayBuffer.
*/
type StartReplayBufferResponse struct {
	requests.ResponseBasic
}

// StartReplayBuffer sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) StartReplayBuffer(
	paramss ...*StartReplayBufferParams,
) (*StartReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartReplayBufferParams{{}}
	}
	params := paramss[0]
	data := &StartReplayBufferResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
