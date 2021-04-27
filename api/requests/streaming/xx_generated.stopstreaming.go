// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StopStreamingParams represents the params body for the "StopStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopStreaming.
*/
type StopStreamingParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of StopStreamingParams
func (o *StopStreamingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StopStreamingParams
func (o *StopStreamingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StopStreamingParams
func (o *StopStreamingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StopStreamingResponse represents the response body for the "StopStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopStreaming.
*/
type StopStreamingResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StopStreamingResponse
func (o *StopStreamingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StopStreamingResponse
func (o *StopStreamingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StopStreamingResponse
func (o *StopStreamingResponse) GetError() string {
	return o.Error
}

// StopStreaming sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) StopStreaming(paramss ...*StopStreamingParams) (*StopStreamingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopStreamingParams{{}}
	}
	params := paramss[0]
	params.RequestType = "StopStreaming"
	data := &StopStreamingResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
