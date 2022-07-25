// This file has been automatically generated. Don't edit it.

package stream

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SendStreamCaptionParams represents the params body for the "SendStreamCaption" request.
Sends CEA-608 caption text over the stream output.
*/
type SendStreamCaptionParams struct {
	requests.ParamsBasic

	// Caption text
	CaptionText string `json:"captionText,omitempty"`
}

// GetSelfName just returns "SendStreamCaption".
func (o *SendStreamCaptionParams) GetSelfName() string {
	return "SendStreamCaption"
}

/*
SendStreamCaptionResponse represents the response body for the "SendStreamCaption" request.
Sends CEA-608 caption text over the stream output.
*/
type SendStreamCaptionResponse struct {
	requests.ResponseBasic
}

// SendStreamCaption sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SendStreamCaption(params *SendStreamCaptionParams) (*SendStreamCaptionResponse, error) {
	data := &SendStreamCaptionResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
