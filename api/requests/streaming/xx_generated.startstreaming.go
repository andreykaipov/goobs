// This file has been automatically generated. Don't edit it.

package streaming

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
StartStreamingParams represents the params body for the "StartStreaming" request.
Start streaming.
Will return an `error` if streaming is already active.
Since 4.1.0.
*/
type StartStreamingParams struct {
	requests.ParamsBasic

	Stream *Stream `json:"stream,omitempty"`
}

type Stream struct {
	// Adds the given object parameters as encoded query string parameters to the 'key' of the RTMP stream. Used to pass
	// data to the RTMP service about the streaming. May be any String, Numeric, or Boolean field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	//
	Settings *typedefs.StreamSettings `json:"settings,omitempty"`

	// If specified ensures the type of stream matches the given type (usually 'rtmp_custom' or 'rtmp_common'). If the
	// currently configured stream type does not match the given stream type, all settings must be specified in the
	// `settings` object or an error will occur when starting the stream.
	Type string `json:"type,omitempty"`
}

// GetSelfName just returns "StartStreaming".
func (o *StartStreamingParams) GetSelfName() string {
	return "StartStreaming"
}

/*
StartStreamingResponse represents the response body for the "StartStreaming" request.
Start streaming.
Will return an `error` if streaming is already active.
Since v4.1.0.
*/
type StartStreamingResponse struct {
	requests.ResponseBasic
}

// StartStreaming sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) StartStreaming(params *StartStreamingParams) (*StartStreamingResponse, error) {
	data := &StartStreamingResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
