// This file has been automatically generated. Don't edit it.

package inputs

/*
SetInputAudioSyncOffsetParams represents the params body for the "SetInputAudioSyncOffset" request.
Sets the audio sync offset of an input.
*/
type SetInputAudioSyncOffsetParams struct {
	// New audio sync offset in milliseconds
	InputAudioSyncOffset float64 `json:"inputAudioSyncOffset,omitempty"`

	// Name of the input to set the audio sync offset of
	InputName string `json:"inputName,omitempty"`
}

/*
SetInputAudioSyncOffsetResponse represents the response body for the "SetInputAudioSyncOffset" request.
Sets the audio sync offset of an input.
*/
type SetInputAudioSyncOffsetResponse struct{}

// SetInputAudioSyncOffset sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputAudioSyncOffset(
	params *SetInputAudioSyncOffsetParams,
) (*SetInputAudioSyncOffsetResponse, error) {
	resp, err := c.SendRequest("SetInputAudioSyncOffset", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetInputAudioSyncOffsetResponse), nil
}
