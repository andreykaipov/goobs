// This file has been automatically generated. Don't edit it.

package outputs

/*
GetReplayBufferStatusParams represents the params body for the "GetReplayBufferStatus" request.
Gets the status of the replay buffer output.
*/
type GetReplayBufferStatusParams struct{}

/*
GetReplayBufferStatusResponse represents the response body for the "GetReplayBufferStatus" request.
Gets the status of the replay buffer output.
*/
type GetReplayBufferStatusResponse struct {
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
	resp, err := c.SendRequest("GetReplayBufferStatus", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetReplayBufferStatusResponse), nil
}
