// This file has been automatically generated. Don't edit it.

package replaybuffer

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SaveReplayBufferParams represents the params body for the "SaveReplayBuffer" request.
Flush and save the contents of the Replay Buffer to disk. This is
basically the same as triggering the "Save Replay Buffer" hotkey.
Will return an `error` if the Replay Buffer is not active.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#SaveReplayBuffer.
*/
type SaveReplayBufferParams struct {
	requests.ParamsBasic
}

// Name just returns "SaveReplayBuffer".
func (o *SaveReplayBufferParams) Name() string {
	return "SaveReplayBuffer"
}

/*
SaveReplayBufferResponse represents the response body for the "SaveReplayBuffer" request.
Flush and save the contents of the Replay Buffer to disk. This is
basically the same as triggering the "Save Replay Buffer" hotkey.
Will return an `error` if the Replay Buffer is not active.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#SaveReplayBuffer.
*/
type SaveReplayBufferResponse struct {
	requests.ResponseBasic
}

// SaveReplayBuffer sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) SaveReplayBuffer(
	paramss ...*SaveReplayBufferParams,
) (*SaveReplayBufferResponse, error) {
	if len(paramss) == 0 {
		paramss = []*SaveReplayBufferParams{{}}
	}
	params := paramss[0]
	data := &SaveReplayBufferResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
