// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the GetInputAudioMonitorType request.
type GetInputAudioMonitorTypeParams struct {
	// Name of the input to get the audio monitor type of
	InputName string `json:"inputName,omitempty"`
}

// Returns the associated request.
func (o *GetInputAudioMonitorTypeParams) GetRequestName() string {
	return "GetInputAudioMonitorType"
}

// Represents the response body for the GetInputAudioMonitorType request.
type GetInputAudioMonitorTypeResponse struct {
	// Audio monitor type
	MonitorType string `json:"monitorType,omitempty"`
}

/*
Gets the audio monitor type of an input.

The available audio monitor types are:

- `OBS_MONITORING_TYPE_NONE`
- `OBS_MONITORING_TYPE_MONITOR_ONLY`
- `OBS_MONITORING_TYPE_MONITOR_AND_OUTPUT`
*/
func (c *Client) GetInputAudioMonitorType(
	params *GetInputAudioMonitorTypeParams,
) (*GetInputAudioMonitorTypeResponse, error) {
	data := &GetInputAudioMonitorTypeResponse{}
	return data, c.SendRequest(params, data)
}
