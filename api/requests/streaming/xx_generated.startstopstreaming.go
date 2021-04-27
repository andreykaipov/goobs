// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartStopStreamingParams represents the params body for the "StartStopStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopStreaming.
*/
type StartStopStreamingParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of StartStopStreamingParams
func (o *StartStopStreamingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StartStopStreamingParams
func (o *StartStopStreamingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StartStopStreamingParams
func (o *StartStopStreamingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StartStopStreamingResponse represents the response body for the "StartStopStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopStreaming.
*/
type StartStopStreamingResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StartStopStreamingResponse
func (o *StartStopStreamingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StartStopStreamingResponse
func (o *StartStopStreamingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StartStopStreamingResponse
func (o *StartStopStreamingResponse) GetError() string {
	return o.Error
}

// StartStopStreaming sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) StartStopStreaming(
	paramss ...*StartStopStreamingParams,
) (*StartStopStreamingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartStopStreamingParams{{}}
	}
	params := paramss[0]
	params.RequestType = "StartStopStreaming"
	data := &StartStopStreamingResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
