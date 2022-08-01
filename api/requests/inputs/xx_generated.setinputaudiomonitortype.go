// This file has been automatically generated. Don't edit it.

package inputs

/*
SetInputAudioMonitorTypeParams represents the params body for the "SetInputAudioMonitorType" request.
Sets the audio monitor type of an input.
*/
type SetInputAudioMonitorTypeParams struct {
	// Name of the input to set the audio monitor type of
	InputName string `json:"inputName,omitempty"`

	// Audio monitor type
	MonitorType string `json:"monitorType,omitempty"`
}

/*
SetInputAudioMonitorTypeResponse represents the response body for the "SetInputAudioMonitorType" request.
Sets the audio monitor type of an input.
*/
type SetInputAudioMonitorTypeResponse struct{}

// SetInputAudioMonitorType sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputAudioMonitorType(
	params *SetInputAudioMonitorTypeParams,
) (*SetInputAudioMonitorTypeResponse, error) {
	resp, err := c.SendRequest("SetInputAudioMonitorType", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetInputAudioMonitorTypeResponse), nil
}
