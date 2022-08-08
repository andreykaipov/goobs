// This file has been automatically generated. Don't edit it.

package stream

// Represents the request body for the SendStreamCaption request.
type SendStreamCaptionParams struct {
	// Caption text
	CaptionText string `json:"captionText,omitempty"`
}

// Returns the associated request.
func (o *SendStreamCaptionParams) GetRequestName() string {
	return "SendStreamCaption"
}

// Represents the response body for the SendStreamCaption request.
type SendStreamCaptionResponse struct{}

// Sends CEA-608 caption text over the stream output.
func (c *Client) SendStreamCaption(params *SendStreamCaptionParams) (*SendStreamCaptionResponse, error) {
	data := &SendStreamCaptionResponse{}
	return data, c.SendRequest(params, data)
}
