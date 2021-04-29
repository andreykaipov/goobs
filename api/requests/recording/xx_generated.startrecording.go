// This file has been automatically generated. Don't edit it.

package recording

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartRecordingParams represents the params body for the "StartRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartRecording.
*/
type StartRecordingParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of StartRecordingParams
func (o *StartRecordingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of StartRecordingParams
func (o *StartRecordingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on StartRecordingParams
func (o *StartRecordingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StartRecordingResponse represents the response body for the "StartRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartRecording.
*/
type StartRecordingResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of StartRecordingResponse
func (o *StartRecordingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of StartRecordingResponse
func (o *StartRecordingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of StartRecordingResponse
func (o *StartRecordingResponse) GetError() string {
	return o.Error
}

// StartRecording sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) StartRecording(paramss ...*StartRecordingParams) (*StartRecordingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartRecordingParams{{}}
	}
	params := paramss[0]
	params.RequestType = "StartRecording"
	data := &StartRecordingResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
