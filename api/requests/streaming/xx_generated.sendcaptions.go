// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SendCaptionsParams represents the params body for the "SendCaptions" request.
Send the provided text as embedded CEA-608 caption data.
As of OBS Studio 23.1, captions are not yet available on Linux.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SendCaptions.
*/
type SendCaptionsParams struct {
	requests.ParamsBasic

	// Captions text
	Text string `json:"text"`
}

// Name just returns "SendCaptions".
func (o *SendCaptionsParams) Name() string {
	return "SendCaptions"
}

/*
SendCaptionsResponse represents the response body for the "SendCaptions" request.
Send the provided text as embedded CEA-608 caption data.
As of OBS Studio 23.1, captions are not yet available on Linux.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SendCaptions.
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
