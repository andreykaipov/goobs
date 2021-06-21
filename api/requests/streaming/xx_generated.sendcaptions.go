// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SendCaptionsParams represents the params body for the "SendCaptions" request.
Send the provided text as embedded CEA-608 caption data.
Since 4.6.0.
*/
type SendCaptionsParams struct {
	requests.ParamsBasic

	// Captions text
	Text string `json:"text,omitempty"`
}

// GetSelfName just returns "SendCaptions".
func (o *SendCaptionsParams) GetSelfName() string {
	return "SendCaptions"
}

/*
SendCaptionsResponse represents the response body for the "SendCaptions" request.
Send the provided text as embedded CEA-608 caption data.
Since v4.6.0.
*/
type SendCaptionsResponse struct {
	requests.ResponseBasic
}

// SendCaptions sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SendCaptions(params *SendCaptionsParams) (*SendCaptionsResponse, error) {
	data := &SendCaptionsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
