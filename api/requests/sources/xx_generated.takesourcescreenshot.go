// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
TakeSourceScreenshotParams represents the params body for the "TakeSourceScreenshot" request.
Takes a picture snapshot of a source and then can either or both:    - Send it over as a Data URI (base64-encoded data) in the response (by specifying `embedPictureFormat` in the request)    - Save it to disk (by specifying `saveToFilePath` in the request)

At least `embedPictureFormat` or `saveToFilePath` must be specified.

Clients can specify `width` and `height` parameters to receive scaled pictures. Aspect ratio is
preserved if only one of these two parameters is specified.

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#TakeSourceScreenshot.
*/
type TakeSourceScreenshotParams struct {
	requests.ParamsBasic

	// Format of the Data URI encoded picture. Can be "png", "jpg", "jpeg" or "bmp" (or any other
	// value supported by Qt's Image module)
	EmbedPictureFormat string `json:"embedPictureFormat"`

	// Screenshot height. Defaults to the source's base height.
	Height int `json:"height"`

	// Full file path (file extension included) where the captured image is to be saved. Can be in a
	// format different from `pictureFormat`. Can be a relative path.
	SaveToFilePath string `json:"saveToFilePath"`

	// Source name
	SourceName string `json:"sourceName"`

	// Screenshot width. Defaults to the source's base width.
	Width int `json:"width"`
}

// Name just returns "TakeSourceScreenshot".
func (o *TakeSourceScreenshotParams) Name() string {
	return "TakeSourceScreenshot"
}

/*
TakeSourceScreenshotResponse represents the response body for the "TakeSourceScreenshot" request.
Takes a picture snapshot of a source and then can either or both:    - Send it over as a Data URI (base64-encoded data) in the response (by specifying `embedPictureFormat` in the request)    - Save it to disk (by specifying `saveToFilePath` in the request)

At least `embedPictureFormat` or `saveToFilePath` must be specified.

Clients can specify `width` and `height` parameters to receive scaled pictures. Aspect ratio is
preserved if only one of these two parameters is specified.

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#TakeSourceScreenshot.
*/
type TakeSourceScreenshotResponse struct {
	requests.ResponseBasic

	// Absolute path to the saved image file (if `saveToFilePath` was specified in the request)
	ImageFile string `json:"imageFile"`

	// Image Data URI (if `embedPictureFormat` was specified in the request)
	Img string `json:"img"`

	// Source name
	SourceName string `json:"sourceName"`
}

// TakeSourceScreenshot sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) TakeSourceScreenshot(
	params *TakeSourceScreenshotParams,
) (*TakeSourceScreenshotResponse, error) {
	data := &TakeSourceScreenshotResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
