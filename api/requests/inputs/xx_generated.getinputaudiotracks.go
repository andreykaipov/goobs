// This file has been automatically generated. Don't edit it.

package inputs

/*
GetInputAudioTracksParams represents the params body for the "GetInputAudioTracks" request.
Gets the enable state of all audio tracks of an input.
*/
type GetInputAudioTracksParams struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`
}

/*
GetInputAudioTracksResponse represents the response body for the "GetInputAudioTracks" request.
Gets the enable state of all audio tracks of an input.
*/
type GetInputAudioTracksResponse struct {
	// Object of audio tracks and associated enable states
	InputAudioTracks interface{} `json:"inputAudioTracks,omitempty"`
}

// GetInputAudioTracks sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputAudioTracks(params *GetInputAudioTracksParams) (*GetInputAudioTracksResponse, error) {
	resp, err := c.SendRequest("GetInputAudioTracks", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetInputAudioTracksResponse), nil
}
