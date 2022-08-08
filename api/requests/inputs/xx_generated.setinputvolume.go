// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the SetInputVolume request.
type SetInputVolumeParams struct {
	// Name of the input to set the volume of
	InputName string `json:"inputName,omitempty"`

	// Volume setting in dB
	InputVolumeDb float64 `json:"inputVolumeDb,omitempty"`

	// Volume setting in mul
	InputVolumeMul float64 `json:"inputVolumeMul,omitempty"`
}

// Returns the associated request.
func (o *SetInputVolumeParams) GetRequestName() string {
	return "SetInputVolume"
}

// Represents the response body for the SetInputVolume request.
type SetInputVolumeResponse struct{}

// Sets the volume setting of an input.
func (c *Client) SetInputVolume(params *SetInputVolumeParams) (*SetInputVolumeResponse, error) {
	data := &SetInputVolumeResponse{}
	return data, c.SendRequest(params, data)
}
