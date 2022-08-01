// This file has been automatically generated. Don't edit it.

package outputs

/*
GetLastReplayBufferReplayParams represents the params body for the "GetLastReplayBufferReplay" request.
Gets the filename of the last replay buffer save file.
*/
type GetLastReplayBufferReplayParams struct{}

/*
GetLastReplayBufferReplayResponse represents the response body for the "GetLastReplayBufferReplay" request.
Gets the filename of the last replay buffer save file.
*/
type GetLastReplayBufferReplayResponse struct {
	// File path
	SavedReplayPath string `json:"savedReplayPath,omitempty"`
}

// GetLastReplayBufferReplay sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetLastReplayBufferReplay(
	paramss ...*GetLastReplayBufferReplayParams,
) (*GetLastReplayBufferReplayResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetLastReplayBufferReplayParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("GetLastReplayBufferReplay", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetLastReplayBufferReplayResponse), nil
}
