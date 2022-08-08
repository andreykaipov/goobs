// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the GetLastReplayBufferReplay request.
type GetLastReplayBufferReplayParams struct{}

// Returns the associated request.
func (o *GetLastReplayBufferReplayParams) GetRequestName() string {
	return "GetLastReplayBufferReplay"
}

// Represents the response body for the GetLastReplayBufferReplay request.
type GetLastReplayBufferReplayResponse struct {
	// File path
	SavedReplayPath string `json:"savedReplayPath,omitempty"`
}

// Gets the filename of the last replay buffer save file.
func (c *Client) GetLastReplayBufferReplay(
	paramss ...*GetLastReplayBufferReplayParams,
) (*GetLastReplayBufferReplayResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetLastReplayBufferReplayParams{{}}
	}
	params := paramss[0]
	data := &GetLastReplayBufferReplayResponse{}
	return data, c.SendRequest(params, data)
}
