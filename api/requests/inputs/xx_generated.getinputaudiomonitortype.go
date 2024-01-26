// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the GetInputAudioMonitorType request.
type GetInputAudioMonitorTypeParams struct {
	// Name of the input to get the audio monitor type of
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input to get the audio monitor type of
	InputUuid *string `json:"inputUuid,omitempty"`
}

func NewGetInputAudioMonitorTypeParams() *GetInputAudioMonitorTypeParams {
	return &GetInputAudioMonitorTypeParams{}
}
func (o *GetInputAudioMonitorTypeParams) WithInputName(x string) *GetInputAudioMonitorTypeParams {
	o.InputName = &x
	return o
}
func (o *GetInputAudioMonitorTypeParams) WithInputUuid(x string) *GetInputAudioMonitorTypeParams {
	o.InputUuid = &x
	return o
}

// Returns the associated request.
func (o *GetInputAudioMonitorTypeParams) GetRequestName() string {
	return "GetInputAudioMonitorType"
}

// Represents the response body for the GetInputAudioMonitorType request.
type GetInputAudioMonitorTypeResponse struct {
	_response

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
	paramss ...*GetInputAudioMonitorTypeParams,
) (*GetInputAudioMonitorTypeResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetInputAudioMonitorTypeParams{{}}
	}
	params := paramss[0]
	data := &GetInputAudioMonitorTypeResponse{}
	return data, c.client.SendRequest(params, data)
}
