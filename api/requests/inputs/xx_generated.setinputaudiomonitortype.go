// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the SetInputAudioMonitorType request.
type SetInputAudioMonitorTypeParams struct {
	// Name of the input to set the audio monitor type of
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input to set the audio monitor type of
	InputUuid *string `json:"inputUuid,omitempty"`

	// Audio monitor type
	MonitorType *string `json:"monitorType,omitempty"`
}

func NewSetInputAudioMonitorTypeParams() *SetInputAudioMonitorTypeParams {
	return &SetInputAudioMonitorTypeParams{}
}
func (o *SetInputAudioMonitorTypeParams) WithInputName(x string) *SetInputAudioMonitorTypeParams {
	o.InputName = &x
	return o
}
func (o *SetInputAudioMonitorTypeParams) WithInputUuid(x string) *SetInputAudioMonitorTypeParams {
	o.InputUuid = &x
	return o
}
func (o *SetInputAudioMonitorTypeParams) WithMonitorType(x string) *SetInputAudioMonitorTypeParams {
	o.MonitorType = &x
	return o
}

// Returns the associated request.
func (o *SetInputAudioMonitorTypeParams) GetRequestName() string {
	return "SetInputAudioMonitorType"
}

// Represents the response body for the SetInputAudioMonitorType request.
type SetInputAudioMonitorTypeResponse struct {
	_response
}

// Sets the audio monitor type of an input.
func (c *Client) SetInputAudioMonitorType(
	params *SetInputAudioMonitorTypeParams,
) (*SetInputAudioMonitorTypeResponse, error) {
	data := &SetInputAudioMonitorTypeResponse{}
	return data, c.client.SendRequest(params, data)
}
