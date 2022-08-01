// This file has been automatically generated. Don't edit it.

package config

/*
GetVideoSettingsParams represents the params body for the "GetVideoSettings" request.
Gets the current video settings.

Note: To get the true FPS value, divide the FPS numerator by the FPS denominator. Example: `60000/1001`
*/
type GetVideoSettingsParams struct{}

/*
GetVideoSettingsResponse represents the response body for the "GetVideoSettings" request.
Gets the current video settings.

Note: To get the true FPS value, divide the FPS numerator by the FPS denominator. Example: `60000/1001`
*/
type GetVideoSettingsResponse struct {
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

// GetVideoSettings sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) GetVideoSettings(paramss ...*GetVideoSettingsParams) (*GetVideoSettingsResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetVideoSettingsParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("GetVideoSettings", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetVideoSettingsResponse), nil
}
