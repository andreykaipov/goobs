// This file has been automatically generated. Don't edit it.

package outputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartReplayBufferParams represents the params body for the "StartReplayBuffer" request.
Starts the replay buffer output.
*/
type StartReplayBufferParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StartReplayBuffer".
func (o *StartReplayBufferParams) GetSelfName() string {
	return "StartReplayBuffer"
}

/*
StartReplayBufferResponse represents the response body for the "StartReplayBuffer" request.
Starts the replay buffer output.
*/
type StartReplayBufferResponse struct {
	requests.ResponseBasic
}

// StartReplayBuffer sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) StartReplayBuffer(paramss ...*StartReplayBufferParams) (*StartReplayBufferResponse, error) {
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
