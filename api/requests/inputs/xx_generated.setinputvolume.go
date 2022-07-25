// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetInputVolumeParams represents the params body for the "SetInputVolume" request.
Sets the volume setting of an input.
*/
type SetInputVolumeParams struct {
	requests.ParamsBasic

	// Name of the input to set the volume of
	InputName string `json:"inputName,omitempty"`

	// Volume setting in dB
	InputVolumeDb float64 `json:"inputVolumeDb,omitempty"`

	// Volume setting in mul
	InputVolumeMul float64 `json:"inputVolumeMul,omitempty"`
}

// GetSelfName just returns "SetInputVolume".
func (o *SetInputVolumeParams) GetSelfName() string {
	return "SetInputVolume"
}

/*
SetInputVolumeResponse represents the response body for the "SetInputVolume" request.
Sets the volume setting of an input.
*/
type SetInputVolumeResponse struct {
	requests.ResponseBasic
}

// SetInputVolume sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetInputVolume(params *SetInputVolumeParams) (*SetInputVolumeResponse, error) {
	data := &SetInputVolumeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
