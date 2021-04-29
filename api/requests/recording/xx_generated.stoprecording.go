// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StopRecordingParams represents the params body for the "StopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopRecording.
*/
type StopRecordingParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of StopRecordingParams
func (o *StopRecordingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StopRecordingParams
func (o *StopRecordingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StopRecordingParams
func (o *StopRecordingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StopRecordingResponse represents the response body for the "StopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopRecording.
*/
type StopRecordingResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StopRecordingResponse
func (o *StopRecordingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StopRecordingResponse
func (o *StopRecordingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StopRecordingResponse
func (o *StopRecordingResponse) GetError() string {
	return o.Error
}

// StopRecording sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) StopRecording(paramss ...*StopRecordingParams) (*StopRecordingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopRecordingParams{{}}
	}
	params := paramss[0]
	params.RequestType = "StopRecording"
	data := &StopRecordingResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
