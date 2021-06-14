// This file has been automatically generated. Don't edit it.

package replaybuffer

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartReplayBufferParams represents the params body for the "StartReplayBuffer" request.
Start recording into the Replay Buffer.
Will return an `error` if the Replay Buffer is already active or if the
"Save Replay Buffer" hotkey is not set in OBS' settings.
Setting this hotkey is mandatory, even when triggering saves only
through obs-websocket.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#StartReplayBuffer.
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
Start recording into the Replay Buffer.
Will return an `error` if the Replay Buffer is already active or if the
"Save Replay Buffer" hotkey is not set in OBS' settings.
Setting this hotkey is mandatory, even when triggering saves only
through obs-websocket.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#StartReplayBuffer.
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
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
