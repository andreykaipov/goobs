// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetVolumeParams represents the params body for the "SetVolume" request.
Set the volume of the specified source. Default request format uses mul, NOT SLIDER PERCENTAGE.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SetVolume.
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

// Name just returns "SetVolume".
func (o *SetVolumeParams) Name() string {
	return "SetVolume"
}

/*
SetVolumeResponse represents the response body for the "SetVolume" request.
Set the volume of the specified source. Default request format uses mul, NOT SLIDER PERCENTAGE.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SetVolume.
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
