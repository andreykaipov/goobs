// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the SetInputVolume request.
type SetInputVolumeParams struct {
	// Name of the input to set the volume of
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input to set the volume of
	InputUuid *string `json:"inputUuid,omitempty"`

	// Volume setting in dB
	InputVolumeDb *float64 `json:"inputVolumeDb,omitempty"`

	// Volume setting in mul
	InputVolumeMul *float64 `json:"inputVolumeMul,omitempty"`
}

func NewSetInputVolumeParams() *SetInputVolumeParams {
	return &SetInputVolumeParams{}
}
func (o *SetInputVolumeParams) WithInputName(x string) *SetInputVolumeParams {
	o.InputName = &x
	return o
}
func (o *SetInputVolumeParams) WithInputUuid(x string) *SetInputVolumeParams {
	o.InputUuid = &x
	return o
}
func (o *SetInputVolumeParams) WithInputVolumeDb(x float64) *SetInputVolumeParams {
	o.InputVolumeDb = &x
	return o
}
func (o *SetInputVolumeParams) WithInputVolumeMul(x float64) *SetInputVolumeParams {
	o.InputVolumeMul = &x
	return o
}

// Returns the associated request.
func (o *SetInputVolumeParams) GetRequestName() string {
	return "SetInputVolume"
}

// Represents the response body for the SetInputVolume request.
type SetInputVolumeResponse struct {
	_response
}

// Sets the volume setting of an input.
func (c *Client) SetInputVolume(paramss ...*SetInputVolumeParams) (*SetInputVolumeResponse, error) {
	if len(paramss) == 0 {
		paramss = []*SetInputVolumeParams{{}}
	}
	params := paramss[0]
	data := &SetInputVolumeResponse{}
	return data, c.client.SendRequest(params, data)
}
