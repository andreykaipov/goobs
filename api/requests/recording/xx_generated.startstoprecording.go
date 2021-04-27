// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartStopRecordingParams represents the params body for the "StartStopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopRecording.
*/
type StartStopRecordingParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of StartStopRecordingParams
func (o *StartStopRecordingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StartStopRecordingParams
func (o *StartStopRecordingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StartStopRecordingParams
func (o *StartStopRecordingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StartStopRecordingResponse represents the response body for the "StartStopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopRecording.
*/
type StartStopRecordingResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StartStopRecordingResponse
func (o *StartStopRecordingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StartStopRecordingResponse
func (o *StartStopRecordingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StartStopRecordingResponse
func (o *StartStopRecordingResponse) GetError() string {
	return o.Error
}

// StartStopRecording sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) StartStopRecording(
	paramss ...*StartStopRecordingParams,
) (*StartStopRecordingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartStopRecordingParams{{}}
	}
	params := paramss[0]
	params.RequestType = "StartStopRecording"
	data := &StartStopRecordingResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
