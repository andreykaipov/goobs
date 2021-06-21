// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetVolumeParams represents the params body for the "GetVolume" request.
Get the volume of the specified source. Default response uses mul format, NOT SLIDER PERCENTAGE.
Since 4.0.0.
*/
type GetVolumeParams struct {
	requests.ParamsBasic

	// Source name.
	Source string `json:"source,omitempty"`

	// Output volume in decibels of attenuation instead of amplitude/mul.
	UseDecibel bool `json:"useDecibel"`
}

// GetSelfName just returns "GetVolume".
func (o *GetVolumeParams) GetSelfName() string {
	return "GetVolume"
}

/*
GetVolumeResponse represents the response body for the "GetVolume" request.
Get the volume of the specified source. Default response uses mul format, NOT SLIDER PERCENTAGE.
Since v4.0.0.
*/
type GetVolumeResponse struct {
	requests.ResponseBasic

	// Indicates whether the source is muted.
	Muted bool `json:"muted"`

	// Source name.
	Name string `json:"name,omitempty"`

	// Volume of the source. Between `0.0` and `20.0` if using mul, under `26.0` if using dB.
	Volume float64 `json:"volume,omitempty"`
}

// GetVolume sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetVolume(params *GetVolumeParams) (*GetVolumeResponse, error) {
	data := &GetVolumeResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
