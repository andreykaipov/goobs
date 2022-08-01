// This file has been automatically generated. Don't edit it.

package inputs

/*
GetInputAudioSyncOffsetParams represents the params body for the "GetInputAudioSyncOffset" request.
Gets the audio sync offset of an input.

Note: The audio sync offset can be negative too!
*/
type GetInputAudioSyncOffsetParams struct {
	// Name of the input to get the audio sync offset of
	InputName string `json:"inputName,omitempty"`
}

/*
GetInputAudioSyncOffsetResponse represents the response body for the "GetInputAudioSyncOffset" request.
Gets the audio sync offset of an input.

Note: The audio sync offset can be negative too!
*/
type GetInputAudioSyncOffsetResponse struct {
	// Audio sync offset in milliseconds
	InputAudioSyncOffset float64 `json:"inputAudioSyncOffset,omitempty"`
}

// GetInputAudioSyncOffset sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputAudioSyncOffset(
	params *GetInputAudioSyncOffsetParams,
) (*GetInputAudioSyncOffsetResponse, error) {
	resp, err := c.SendRequest("GetInputAudioSyncOffset", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetInputAudioSyncOffsetResponse), nil
}
