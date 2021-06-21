// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
TakeSourceScreenshotParams represents the params body for the "TakeSourceScreenshot" request.
Takes a picture snapshot of a source and then can either or both: 	- Send it over as a Data URI (base64-encoded data) in the response (by specifying `embedPictureFormat` in the request) 	- Save it to disk (by specifying `saveToFilePath` in the request)

At least `embedPictureFormat` or `saveToFilePath` must be specified.

Clients can specify `width` and `height` parameters to receive scaled pictures. Aspect ratio is
preserved if only one of these two parameters is specified.
Since 4.6.0.
*/
type TakeSourceScreenshotParams struct {
	requests.ParamsBasic

	// Compression ratio between -1 and 100 to write the image with. -1 is automatic, 1 is smallest file/most
	// compression, 100 is largest file/least compression. Varies with image type.
	CompressionQuality int `json:"compressionQuality,omitempty"`

	// Format of the Data URI encoded picture. Can be "png", "jpg", "jpeg" or "bmp" (or any other value supported by
	// Qt's Image module)
	EmbedPictureFormat string `json:"embedPictureFormat,omitempty"`

	// Format to save the image file as (one of the values provided in the `supported-image-export-formats` response
	// field of `GetVersion`). If not specified, tries to guess based on file extension.
	FileFormat string `json:"fileFormat,omitempty"`

	// Screenshot height. Defaults to the source's base height.
	Height int `json:"height,omitempty"`

	// Full file path (file extension included) where the captured image is to be saved. Can be in a format different
	// from `pictureFormat`. Can be a relative path.
	SaveToFilePath string `json:"saveToFilePath,omitempty"`

	// Source name. Note: Since scenes are also sources, you can also provide a scene name. If not provided, the
	// currently active scene is used.
	SourceName string `json:"sourceName,omitempty"`

	// Screenshot width. Defaults to the source's base width.
	Width int `json:"width,omitempty"`
}

// GetSelfName just returns "TakeSourceScreenshot".
func (o *TakeSourceScreenshotParams) GetSelfName() string {
	return "TakeSourceScreenshot"
}

/*
TakeSourceScreenshotResponse represents the response body for the "TakeSourceScreenshot" request.
Takes a picture snapshot of a source and then can either or both: 	- Send it over as a Data URI (base64-encoded data) in the response (by specifying `embedPictureFormat` in the request) 	- Save it to disk (by specifying `saveToFilePath` in the request)

At least `embedPictureFormat` or `saveToFilePath` must be specified.

Clients can specify `width` and `height` parameters to receive scaled pictures. Aspect ratio is
preserved if only one of these two parameters is specified.
Since v4.6.0.
*/
type TakeSourceScreenshotResponse struct {
	requests.ResponseBasic

	// Absolute path to the saved image file (if `saveToFilePath` was specified in the request)
	ImageFile string `json:"imageFile,omitempty"`

	// Image Data URI (if `embedPictureFormat` was specified in the request)
	Img string `json:"img,omitempty"`

	// Source name
	SourceName string `json:"sourceName,omitempty"`
}

// TakeSourceScreenshot sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) TakeSourceScreenshot(params *TakeSourceScreenshotParams) (*TakeSourceScreenshotResponse, error) {
	data := &TakeSourceScreenshotResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
