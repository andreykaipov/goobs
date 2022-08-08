// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the GetInputVolume request.
type GetInputVolumeParams struct {
	// Name of the input to get the volume of
	InputName string `json:"inputName,omitempty"`
}

// Returns the associated request.
func (o *GetInputVolumeParams) GetRequestName() string {
	return "GetInputVolume"
}

// Represents the response body for the GetInputVolume request.
type GetInputVolumeResponse struct {
	// Volume setting in dB
	InputVolumeDb float64 `json:"inputVolumeDb,omitempty"`

	// Volume setting in mul
	InputVolumeMul float64 `json:"inputVolumeMul,omitempty"`
}

// Gets the current volume setting of an input.
func (c *Client) GetInputVolume(params *GetInputVolumeParams) (*GetInputVolumeResponse, error) {
	data := &GetInputVolumeResponse{}
	return data, c.SendRequest(params, data)
}
