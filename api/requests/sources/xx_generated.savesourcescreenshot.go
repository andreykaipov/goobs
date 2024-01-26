// This file has been automatically generated. Don't edit it.

package sources

// Represents the request body for the SaveSourceScreenshot request.
type SaveSourceScreenshotParams struct {
	// Compression quality to use. 0 for high compression, 100 for uncompressed. -1 to use "default" (whatever that
	// means, idk)
	ImageCompressionQuality *float64 `json:"imageCompressionQuality,omitempty"`

	// Path to save the screenshot file to. Eg. `C:\Users\user\Desktop\screenshot.png`
	ImageFilePath *string `json:"imageFilePath,omitempty"`

	// Image compression format to use. Use `GetVersion` to get compatible image formats
	ImageFormat *string `json:"imageFormat,omitempty"`

	// Height to scale the screenshot to
	ImageHeight *float64 `json:"imageHeight,omitempty"`

	// Width to scale the screenshot to
	ImageWidth *float64 `json:"imageWidth,omitempty"`

	// Name of the source to take a screenshot of
	SourceName *string `json:"sourceName,omitempty"`

	// UUID of the source to take a screenshot of
	SourceUuid *string `json:"sourceUuid,omitempty"`
}

func NewSaveSourceScreenshotParams() *SaveSourceScreenshotParams {
	return &SaveSourceScreenshotParams{}
}
func (o *SaveSourceScreenshotParams) WithImageCompressionQuality(x float64) *SaveSourceScreenshotParams {
	o.ImageCompressionQuality = &x
	return o
}
func (o *SaveSourceScreenshotParams) WithImageFilePath(x string) *SaveSourceScreenshotParams {
	o.ImageFilePath = &x
	return o
}
func (o *SaveSourceScreenshotParams) WithImageFormat(x string) *SaveSourceScreenshotParams {
	o.ImageFormat = &x
	return o
}
func (o *SaveSourceScreenshotParams) WithImageHeight(x float64) *SaveSourceScreenshotParams {
	o.ImageHeight = &x
	return o
}
func (o *SaveSourceScreenshotParams) WithImageWidth(x float64) *SaveSourceScreenshotParams {
	o.ImageWidth = &x
	return o
}
func (o *SaveSourceScreenshotParams) WithSourceName(x string) *SaveSourceScreenshotParams {
	o.SourceName = &x
	return o
}
func (o *SaveSourceScreenshotParams) WithSourceUuid(x string) *SaveSourceScreenshotParams {
	o.SourceUuid = &x
	return o
}

// Returns the associated request.
func (o *SaveSourceScreenshotParams) GetRequestName() string {
	return "SaveSourceScreenshot"
}

// Represents the response body for the SaveSourceScreenshot request.
type SaveSourceScreenshotResponse struct {
	_response
}

/*
Saves a screenshot of a source to the filesystem.

The `imageWidth` and `imageHeight` parameters are treated as "scale to inner", meaning the smallest ratio will be used and the aspect ratio of the original resolution is kept.
If `imageWidth` and `imageHeight` are not specified, the compressed image will use the full resolution of the source.

**Compatible with inputs and scenes.**
*/
func (c *Client) SaveSourceScreenshot(params *SaveSourceScreenshotParams) (*SaveSourceScreenshotResponse, error) {
	data := &SaveSourceScreenshotResponse{}
	return data, c.client.SendRequest(params, data)
}
