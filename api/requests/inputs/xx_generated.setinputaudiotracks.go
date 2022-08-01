// This file has been automatically generated. Don't edit it.

package inputs

/*
SetInputAudioTracksParams represents the params body for the "SetInputAudioTracks" request.
Sets the enable state of audio tracks of an input.
*/
type SetInputAudioTracksParams struct {
	// Track settings to apply
	InputAudioTracks interface{} `json:"inputAudioTracks,omitempty"`

	// Name of the input
	InputName string `json:"inputName,omitempty"`
}

/*
SetInputAudioTracksResponse represents the response body for the "SetInputAudioTracks" request.
Sets the enable state of audio tracks of an input.
*/
type SetInputAudioTracksResponse struct{}

// SetInputAudioTracks sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputAudioTracks(params *SetInputAudioTracksParams) (*SetInputAudioTracksResponse, error) {
	resp, err := c.SendRequest("SetInputAudioTracks", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetInputAudioTracksResponse), nil
}
