// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetVideoInfoParams represents the params body for the "GetVideoInfo" request.
Get basic OBS video information
Since 4.6.0.
*/
type GetVideoInfoParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetVideoInfo".
func (o *GetVideoInfoParams) GetSelfName() string {
	return "GetVideoInfo"
}

/*
GetVideoInfoResponse represents the response body for the "GetVideoInfo" request.
Get basic OBS video information
Since v4.6.0.
*/
type GetVideoInfoResponse struct {
	requests.ResponseBasic

	// Base (canvas) height
	BaseHeight int `json:"baseHeight,omitempty"`

	// Base (canvas) width
	BaseWidth int `json:"baseWidth,omitempty"`

	// Color range (full or partial)
	ColorRange string `json:"colorRange,omitempty"`

	// Color space for YUV
	ColorSpace string `json:"colorSpace,omitempty"`

	// Frames rendered per second
	Fps float64 `json:"fps,omitempty"`

	// Output height
	OutputHeight int `json:"outputHeight,omitempty"`

	// Output width
	OutputWidth int `json:"outputWidth,omitempty"`

	// Scaling method used if output size differs from base size
	ScaleType string `json:"scaleType,omitempty"`

	// Video color format
	VideoFormat string `json:"videoFormat,omitempty"`
}

// GetVideoInfo sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) GetVideoInfo(paramss ...*GetVideoInfoParams) (*GetVideoInfoResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetVideoInfoParams{{}}
	}
	params := paramss[0]
	data := &GetVideoInfoResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
