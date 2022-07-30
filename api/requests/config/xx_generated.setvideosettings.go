// This file has been automatically generated. Don't edit it.

package config

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetVideoSettingsParams represents the params body for the "SetVideoSettings" request.
Sets the current video settings.

Note: Fields must be specified in pairs. For example, you cannot set only `baseWidth` without needing to specify `baseHeight`.
*/
type SetVideoSettingsParams struct {
	requests.ParamsBasic

	// Height of the base (canvas) resolution in pixels
	BaseHeight float64 `json:"baseHeight,omitempty"`

	// Width of the base (canvas) resolution in pixels
	BaseWidth float64 `json:"baseWidth,omitempty"`

	// Denominator of the fractional FPS value
	FpsDenominator float64 `json:"fpsDenominator,omitempty"`

	// Numerator of the fractional FPS value
	FpsNumerator float64 `json:"fpsNumerator,omitempty"`

	// Height of the output resolution in pixels
	OutputHeight float64 `json:"outputHeight,omitempty"`

	// Width of the output resolution in pixels
	OutputWidth float64 `json:"outputWidth,omitempty"`
}

// GetSelfName just returns "SetVideoSettings".
func (o *SetVideoSettingsParams) GetSelfName() string {
	return "SetVideoSettings"
}

/*
SetVideoSettingsResponse represents the response body for the "SetVideoSettings" request.
Sets the current video settings.

Note: Fields must be specified in pairs. For example, you cannot set only `baseWidth` without needing to specify `baseHeight`.
*/
type SetVideoSettingsResponse struct {
	requests.ResponseBasic
}

// SetVideoSettings sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) SetVideoSettings(paramss ...*SetVideoSettingsParams) (*SetVideoSettingsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*SetVideoSettingsParams{{}}
	}
	params := paramss[0]
	data := &SetVideoSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
