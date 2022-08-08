// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the SetInputAudioMonitorType request.
type SetInputAudioMonitorTypeParams struct {
	// Name of the input to set the audio monitor type of
	InputName string `json:"inputName,omitempty"`

	// Audio monitor type
	MonitorType string `json:"monitorType,omitempty"`
}

// Returns the associated request.
func (o *SetInputAudioMonitorTypeParams) GetRequestName() string {
	return "SetInputAudioMonitorType"
}

// Represents the response body for the SetInputAudioMonitorType request.
type SetInputAudioMonitorTypeResponse struct{}

// Sets the audio monitor type of an input.
func (c *Client) SetInputAudioMonitorType(
	params *SetInputAudioMonitorTypeParams,
) (*SetInputAudioMonitorTypeResponse, error) {
	data := &SetInputAudioMonitorTypeResponse{}
	return data, c.SendRequest(params, data)
}
