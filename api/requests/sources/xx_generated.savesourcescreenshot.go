// This file has been automatically generated. Don't edit it.

package sources

// Represents the request body for the SaveSourceScreenshot request.
type SaveSourceScreenshotParams struct {
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

// Returns the associated request.
func (o *SaveSourceScreenshotParams) GetRequestName() string {
	return "SaveSourceScreenshot"
}

// Represents the response body for the SaveSourceScreenshot request.
type SaveSourceScreenshotResponse struct {
	// Base64-encoded screenshot
	ImageData string `json:"imageData,omitempty"`
}

/*
Saves a screenshot of a source to the filesystem.

The `imageWidth` and `imageHeight` parameters are treated as "scale to inner", meaning the smallest ratio will be used and the aspect ratio of the original resolution is kept.
If `imageWidth` and `imageHeight` are not specified, the compressed image will use the full resolution of the source.

**Compatible with inputs and scenes.**
*/
func (c *Client) SaveSourceScreenshot(params *SaveSourceScreenshotParams) (*SaveSourceScreenshotResponse, error) {
	data := &SaveSourceScreenshotResponse{}
	return data, c.SendRequest(params, data)
}
