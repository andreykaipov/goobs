// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetVolumeParams represents the params body for the "GetVolume" request.
Get the volume of the specified source. Default response uses mul format, NOT SLIDER PERCENTAGE.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetVolume.
*/
type GetVolumeParams struct {
	requests.ParamsBasic

	// Source name.
	Source string `json:"source"`

	// Output volume in decibels of attenuation instead of amplitude/mul.
	UseDecibel bool `json:"useDecibel"`
}

// Name just returns "GetVolume".
func (o *GetVolumeParams) Name() string {
	return "GetVolume"
}

/*
GetVolumeResponse represents the response body for the "GetVolume" request.
Get the volume of the specified source. Default response uses mul format, NOT SLIDER PERCENTAGE.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetVolume.
*/
type GetVolumeResponse struct {
	requests.ResponseBasic

	// Indicates whether the source is muted.
	Muted bool `json:"muted"`

	// Source name.
	Name string `json:"name"`

	// Volume of the source. Between `0.0` and `1.0` if using mul, under `0.0` if using dB (since it
	// is attenuating).
	Volume float64 `json:"volume"`
}

// GetVolume sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetVolume(params *GetVolumeParams) (*GetVolumeResponse, error) {
	data := &GetVolumeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
