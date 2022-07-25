// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SaveSourceScreenshotParams represents the params body for the "SaveSourceScreenshot" request.
Saves a screenshot of a source to the filesystem.

The `imageWidth` and `imageHeight` parameters are treated as "scale to inner", meaning the smallest ratio will be used and the aspect ratio of the original resolution is kept.
If `imageWidth` and `imageHeight` are not specified, the compressed image will use the full resolution of the source.

**Compatible with inputs and scenes.**
*/
type SaveSourceScreenshotParams struct {
	requests.ParamsBasic

	// Compression quality to use. 0 for high compression, 100 for uncompressed. -1 to use "default" (whatever that
	// means, idk)
	ImageCompressionQuality float64 `json:"imageCompressionQuality,omitempty"`

	// Path to save the screenshot file to. Eg. `C:\Users\user\Desktop\screenshot.png`
	ImageFilePath string `json:"imageFilePath,omitempty"`

	// Image compression format to use. Use `GetVersion` to get compatible image formats
	ImageFormat string `json:"imageFormat,omitempty"`

	// Height to scale the screenshot to
	ImageHeight float64 `json:"imageHeight,omitempty"`

	// Width to scale the screenshot to
	ImageWidth float64 `json:"imageWidth,omitempty"`

	// Name of the source to take a screenshot of
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "SaveSourceScreenshot".
func (o *SaveSourceScreenshotParams) GetSelfName() string {
	return "SaveSourceScreenshot"
}

/*
SaveSourceScreenshotResponse represents the response body for the "SaveSourceScreenshot" request.
Saves a screenshot of a source to the filesystem.

The `imageWidth` and `imageHeight` parameters are treated as "scale to inner", meaning the smallest ratio will be used and the aspect ratio of the original resolution is kept.
If `imageWidth` and `imageHeight` are not specified, the compressed image will use the full resolution of the source.

**Compatible with inputs and scenes.**
*/
type SaveSourceScreenshotResponse struct {
	requests.ResponseBasic

	// Base64-encoded screenshot
	ImageData string `json:"imageData,omitempty"`
}

// SaveSourceScreenshot sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SaveSourceScreenshot(params *SaveSourceScreenshotParams) (*SaveSourceScreenshotResponse, error) {
	data := &SaveSourceScreenshotResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
