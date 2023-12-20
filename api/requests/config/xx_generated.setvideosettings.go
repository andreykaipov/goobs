// This file has been automatically generated. Don't edit it.

package config

// Represents the request body for the SetVideoSettings request.
type SetVideoSettingsParams struct {
	// Height of the base (canvas) resolution in pixels
	BaseHeight *float64 `json:"baseHeight,omitempty"`

	// Width of the base (canvas) resolution in pixels
	BaseWidth *float64 `json:"baseWidth,omitempty"`

	// Denominator of the fractional FPS value
	FpsDenominator *float64 `json:"fpsDenominator,omitempty"`

	// Numerator of the fractional FPS value
	FpsNumerator *float64 `json:"fpsNumerator,omitempty"`

	// Height of the output resolution in pixels
	OutputHeight *float64 `json:"outputHeight,omitempty"`

	// Width of the output resolution in pixels
	OutputWidth *float64 `json:"outputWidth,omitempty"`
}

func NewSetVideoSettingsParams() *SetVideoSettingsParams {
	return &SetVideoSettingsParams{}
}
func (o *SetVideoSettingsParams) WithBaseHeight(x float64) *SetVideoSettingsParams {
	o.BaseHeight = &x
	return o
}
func (o *SetVideoSettingsParams) WithBaseWidth(x float64) *SetVideoSettingsParams {
	o.BaseWidth = &x
	return o
}
func (o *SetVideoSettingsParams) WithFpsDenominator(x float64) *SetVideoSettingsParams {
	o.FpsDenominator = &x
	return o
}
func (o *SetVideoSettingsParams) WithFpsNumerator(x float64) *SetVideoSettingsParams {
	o.FpsNumerator = &x
	return o
}
func (o *SetVideoSettingsParams) WithOutputHeight(x float64) *SetVideoSettingsParams {
	o.OutputHeight = &x
	return o
}
func (o *SetVideoSettingsParams) WithOutputWidth(x float64) *SetVideoSettingsParams {
	o.OutputWidth = &x
	return o
}

// Returns the associated request.
func (o *SetVideoSettingsParams) GetRequestName() string {
	return "SetVideoSettings"
}

// Represents the response body for the SetVideoSettings request.
type SetVideoSettingsResponse struct {
	_response
}

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
	return data, c.client.SendRequest(params, data)
}
