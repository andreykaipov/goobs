// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the GetInputVolume request.
type GetInputVolumeParams struct {
	// Name of the input to get the volume of
	InputName *string `json:"inputName,omitempty"`

	// UUID of the input to get the volume of
	InputUuid *string `json:"inputUuid,omitempty"`
}

func NewGetInputVolumeParams() *GetInputVolumeParams {
	return &GetInputVolumeParams{}
}
func (o *GetInputVolumeParams) WithInputName(x string) *GetInputVolumeParams {
	o.InputName = &x
	return o
}
func (o *GetInputVolumeParams) WithInputUuid(x string) *GetInputVolumeParams {
	o.InputUuid = &x
	return o
}

// Returns the associated request.
func (o *GetInputVolumeParams) GetRequestName() string {
	return "GetInputVolume"
}

// Represents the response body for the GetInputVolume request.
type GetInputVolumeResponse struct {
	_response

	// Volume setting in dB
	InputVolumeDb float64 `json:"inputVolumeDb,omitempty"`

	// Volume setting in mul
	InputVolumeMul float64 `json:"inputVolumeMul,omitempty"`
}

// Gets the current volume setting of an input.
func (c *Client) GetInputVolume(paramss ...*GetInputVolumeParams) (*GetInputVolumeResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetInputVolumeParams{{}}
	}
	params := paramss[0]
	data := &GetInputVolumeResponse{}
	return data, c.client.SendRequest(params, data)
}
