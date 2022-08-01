// This file has been automatically generated. Don't edit it.

package inputs

/*
GetInputVolumeParams represents the params body for the "GetInputVolume" request.
Gets the current volume setting of an input.
*/
type GetInputVolumeParams struct {
	// Name of the input to get the volume of
	InputName string `json:"inputName,omitempty"`
}

/*
GetInputVolumeResponse represents the response body for the "GetInputVolume" request.
Gets the current volume setting of an input.
*/
type GetInputVolumeResponse struct {
	// Volume setting in dB
	InputVolumeDb float64 `json:"inputVolumeDb,omitempty"`

	// Volume setting in mul
	InputVolumeMul float64 `json:"inputVolumeMul,omitempty"`
}

// GetInputVolume sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputVolume(params *GetInputVolumeParams) (*GetInputVolumeResponse, error) {
	resp, err := c.SendRequest("GetInputVolume", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetInputVolumeResponse), nil
}
