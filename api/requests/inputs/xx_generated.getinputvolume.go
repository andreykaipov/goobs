// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetInputVolumeParams represents the params body for the "GetInputVolume" request.
Gets the current volume setting of an input.
*/
type GetInputVolumeParams struct {
	requests.ParamsBasic

	// Name of the input to get the volume of
	InputName string `json:"inputName,omitempty"`
}

// GetSelfName just returns "GetInputVolume".
func (o *GetInputVolumeParams) GetSelfName() string {
	return "GetInputVolume"
}

/*
GetInputVolumeResponse represents the response body for the "GetInputVolume" request.
Gets the current volume setting of an input.
*/
type GetInputVolumeResponse struct {
	requests.ResponseBasic

	// Volume setting in dB
	InputVolumeDb float64 `json:"inputVolumeDb,omitempty"`

	// Volume setting in mul
	InputVolumeMul float64 `json:"inputVolumeMul,omitempty"`
}

// GetInputVolume sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputVolume(params *GetInputVolumeParams) (*GetInputVolumeResponse, error) {
	data := &GetInputVolumeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
