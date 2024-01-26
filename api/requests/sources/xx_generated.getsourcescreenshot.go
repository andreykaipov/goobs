// This file has been automatically generated. Don't edit it.

package sources

// Represents the request body for the GetSourceScreenshot request.
type GetSourceScreenshotParams struct {
	// Compression quality to use. 0 for high compression, 100 for uncompressed. -1 to use "default" (whatever that
	// means, idk)
	ImageCompressionQuality *float64 `json:"imageCompressionQuality,omitempty"`

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

func NewGetSourceScreenshotParams() *GetSourceScreenshotParams {
	return &GetSourceScreenshotParams{}
}
func (o *GetSourceScreenshotParams) WithImageCompressionQuality(x float64) *GetSourceScreenshotParams {
	o.ImageCompressionQuality = &x
	return o
}
func (o *GetSourceScreenshotParams) WithImageFormat(x string) *GetSourceScreenshotParams {
	o.ImageFormat = &x
	return o
}
func (o *GetSourceScreenshotParams) WithImageHeight(x float64) *GetSourceScreenshotParams {
	o.ImageHeight = &x
	return o
}
func (o *GetSourceScreenshotParams) WithImageWidth(x float64) *GetSourceScreenshotParams {
	o.ImageWidth = &x
	return o
}
func (o *GetSourceScreenshotParams) WithSourceName(x string) *GetSourceScreenshotParams {
	o.SourceName = &x
	return o
}
func (o *GetSourceScreenshotParams) WithSourceUuid(x string) *GetSourceScreenshotParams {
	o.SourceUuid = &x
	return o
}

// Returns the associated request.
func (o *GetSourceScreenshotParams) GetRequestName() string {
	return "GetSourceScreenshot"
}

// Represents the response body for the GetSourceScreenshot request.
type GetSourceScreenshotResponse struct {
	_response

	// Base64-encoded screenshot
	ImageData string `json:"imageData,omitempty"`
}

/*
Gets a Base64-encoded screenshot of a source.

The `imageWidth` and `imageHeight` parameters are treated as "scale to inner", meaning the smallest ratio will be used and the aspect ratio of the original resolution is kept.
If `imageWidth` and `imageHeight` are not specified, the compressed image will use the full resolution of the source.

**Compatible with inputs and scenes.**
*/
func (c *Client) GetSourceScreenshot(params *GetSourceScreenshotParams) (*GetSourceScreenshotResponse, error) {
	data := &GetSourceScreenshotResponse{}
	return data, c.client.SendRequest(params, data)
}
