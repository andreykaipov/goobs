// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartStreamingParams represents the params body for the "StartStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStreaming.
*/
type StartStreamingParams struct {
	requests.Params

	Stream struct {
		// Adds the given object parameters as encoded query string parameters to the 'key' of the
		// RTMP stream. Used to pass data to the RTMP service about the streaming. May be any
		// String, Numeric, or Boolean field.
		Metadata map[string]interface{} `json:"metadata"`

		Settings struct {
			// The publish key of the stream.
			Key string `json:"key"`

			// If authentication is enabled, the password for the streaming server. Ignored if
			// `use-auth` is not set to `true`.
			Password string `json:"password"`

			// The publish URL.
			Server string `json:"server"`

			// Indicates whether authentication should be used when connecting to the streaming
			// server.
			UseAuth bool `json:"use-auth"`

			// If authentication is enabled, the username for the streaming server. Ignored if
			// `use-auth` is not set to `true`.
			Username string `json:"username"`
		} `json:"settings"`

		// If specified ensures the type of stream matches the given type (usually 'rtmp_custom' or
		// 'rtmp_common'). If the currently configured stream type does not match the given stream
		// type, all settings must be specified in the `settings` object or an error will occur when
		// starting the stream.
		Type string `json:"type"`
	} `json:"stream"`
}

// GetRequestType returns the RequestType of StartStreamingParams
func (o *StartStreamingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StartStreamingParams
func (o *StartStreamingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StartStreamingParams
func (o *StartStreamingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StartStreamingResponse represents the response body for the "StartStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStreaming.
*/
type StartStreamingResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StartStreamingResponse
func (o *StartStreamingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StartStreamingResponse
func (o *StartStreamingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StartStreamingResponse
func (o *StartStreamingResponse) GetError() string {
	return o.Error
}

// StartStreaming sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) StartStreaming(params *StartStreamingParams) (*StartStreamingResponse, error) {
	params.RequestType = "StartStreaming"
	data := &StartStreamingResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
