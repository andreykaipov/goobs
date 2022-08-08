// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the GetReplayBufferStatus request.
type GetReplayBufferStatusParams struct{}

// Returns the associated request.
func (o *GetReplayBufferStatusParams) GetRequestName() string {
	return "GetReplayBufferStatus"
}

// Represents the response body for the GetReplayBufferStatus request.
type GetReplayBufferStatusResponse struct {
	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`
}

// Gets the status of the replay buffer output.
func (c *Client) GetReplayBufferStatus(
	paramss ...*GetReplayBufferStatusParams,
) (*GetReplayBufferStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetReplayBufferStatusParams{{}}
	}
	params := paramss[0]
	data := &GetReplayBufferStatusResponse{}
	return data, c.SendRequest(params, data)
}
