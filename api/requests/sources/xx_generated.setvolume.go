// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetVolumeParams represents the params body for the "SetVolume" request.
Set the volume of the specified source. Default request format uses mul, NOT SLIDER PERCENTAGE.
Since 4.0.0.
*/
type SetVolumeParams struct {
	requests.ParamsBasic

	// Source name.
	Source string `json:"source"`

	// Interperet `volume` data as decibels instead of amplitude/mul.
	UseDecibel bool `json:"useDecibel"`

	// Desired volume. Must be between `0.0` and `1.0` for mul, and under 0.0 for dB. Note: OBS will
	// interpret dB values under -100.0 as Inf.
	Volume float64 `json:"volume"`
}

// GetSelfName just returns "SetVolume".
func (o *SetVolumeParams) GetSelfName() string {
	return "SetVolume"
}

/*
SetVolumeResponse represents the response body for the "SetVolume" request.
Set the volume of the specified source. Default request format uses mul, NOT SLIDER PERCENTAGE.
Since v4.0.0.
*/
type SetVolumeResponse struct {
	requests.ResponseBasic
}

// SetVolume sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetVolume(params *SetVolumeParams) (*SetVolumeResponse, error) {
	data := &SetVolumeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
