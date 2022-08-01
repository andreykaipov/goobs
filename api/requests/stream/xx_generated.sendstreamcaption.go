// This file has been automatically generated. Don't edit it.

package stream

/*
SendStreamCaptionParams represents the params body for the "SendStreamCaption" request.
Sends CEA-608 caption text over the stream output.
*/
type SendStreamCaptionParams struct {
	// Caption text
	CaptionText string `json:"captionText,omitempty"`
}

/*
SendStreamCaptionResponse represents the response body for the "SendStreamCaption" request.
Sends CEA-608 caption text over the stream output.
*/
type SendStreamCaptionResponse struct{}

// SendStreamCaption sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SendStreamCaption(params *SendStreamCaptionParams) (*SendStreamCaptionResponse, error) {
	resp, err := c.SendRequest("SendStreamCaption", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SendStreamCaptionResponse), nil
}
