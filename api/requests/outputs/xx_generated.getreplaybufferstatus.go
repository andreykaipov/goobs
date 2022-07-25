// This file has been automatically generated. Don't edit it.

package outputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetReplayBufferStatusParams represents the params body for the "GetReplayBufferStatus" request.
Gets the status of the replay buffer output.
*/
type GetReplayBufferStatusParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetReplayBufferStatus".
func (o *GetReplayBufferStatusParams) GetSelfName() string {
	return "GetReplayBufferStatus"
}

/*
GetReplayBufferStatusResponse represents the response body for the "GetReplayBufferStatus" request.
Gets the status of the replay buffer output.
*/
type GetReplayBufferStatusResponse struct {
	requests.ResponseBasic

	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`
}

// GetReplayBufferStatus sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetReplayBufferStatus(
	paramss ...*GetReplayBufferStatusParams,
) (*GetReplayBufferStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetReplayBufferStatusParams{{}}
	}
	params := paramss[0]
	data := &GetReplayBufferStatusResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
