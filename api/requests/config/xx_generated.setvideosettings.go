// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the SetVideoSettings request.
type SetVideoSettingsParams struct {
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

// Returns the associated request.
func (o *SetVideoSettingsParams) GetRequestName() string {
	return "SetVideoSettings"
}

// Represents the response body for the SetVideoSettings request.
type SetVideoSettingsResponse struct{}

/*
Sets the current video settings.

Note: Fields must be specified in pairs. For example, you cannot set only `baseWidth` without needing to specify `baseHeight`.
*/
func (c *Client) SetVideoSettings(paramss ...*SetVideoSettingsParams) (*SetVideoSettingsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*SetVideoSettingsParams{{}}
	}
	params := paramss[0]
	data := &SetVideoSettingsResponse{}
	return data, c.SendRequest(params, data)
}
