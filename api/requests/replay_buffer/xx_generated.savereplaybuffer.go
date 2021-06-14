// This file has been automatically generated. Don't edit it.

package replaybuffer

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SaveReplayBufferParams represents the params body for the "SaveReplayBuffer" request.
Flush and save the contents of the Replay Buffer to disk. This is
basically the same as triggering the "Save Replay Buffer" hotkey.
Will return an `error` if the Replay Buffer is not active.
Since 4.2.0.
*/
type SaveReplayBufferParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "SaveReplayBuffer".
func (o *SaveReplayBufferParams) GetSelfName() string {
	return "SaveReplayBuffer"
}

/*
SaveReplayBufferResponse represents the response body for the "SaveReplayBuffer" request.
Flush and save the contents of the Replay Buffer to disk. This is
basically the same as triggering the "Save Replay Buffer" hotkey.
Will return an `error` if the Replay Buffer is not active.
Since v4.2.0.
*/
type SaveReplayBufferResponse struct {
	requests.ResponseBasic
}

// SaveReplayBuffer sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) SaveReplayBuffer(paramss ...*SaveReplayBufferParams) (*SaveReplayBufferResponse, error) {
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
