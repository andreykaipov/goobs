// This file has been automatically generated. Don't edit it.

package inputs

/*
SetInputVolumeParams represents the params body for the "SetInputVolume" request.
Sets the volume setting of an input.
*/
type SetInputVolumeParams struct {
	// Name of the input to set the volume of
	InputName string `json:"inputName,omitempty"`

	// Volume setting in dB
	InputVolumeDb float64 `json:"inputVolumeDb,omitempty"`

	// Volume setting in mul
	InputVolumeMul float64 `json:"inputVolumeMul,omitempty"`
}

/*
SetInputVolumeResponse represents the response body for the "SetInputVolume" request.
Sets the volume setting of an input.
*/
type SetInputVolumeResponse struct{}

// SetInputVolume sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputVolume(params *SetInputVolumeParams) (*SetInputVolumeResponse, error) {
	resp, err := c.SendRequest("SetInputVolume", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetInputVolumeResponse), nil
}
