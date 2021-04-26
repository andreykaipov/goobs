// This file has been automatically generated. Don't edit it.

package recording

import api "github.com/andreykaipov/goobs/api"

/*
StartStopRecordingParams represents the params body for the "StartStopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopRecording.
*/
type StartStopRecordingParams struct {
	api.Params
}

func (o *StartStopRecordingParams) GetRequestType() string {
	return o.RequestType
}
func (o *StartStopRecordingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StartStopRecordingResponse represents the response body for the "StartStopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopRecording.
*/
type StartStopRecordingResponse struct {
	api.Response
}

func (o *StartStopRecordingResponse) GetMessageID() string {
	return o.MessageID
}
func (o *StartStopRecordingResponse) GetStatus() string {
	return o.Status
}
func (o *StartStopRecordingResponse) GetError() string {
	return o.Error
}

/*
StartRecordingParams represents the params body for the "StartRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartRecording.
*/
type StartRecordingParams struct {
	api.Params
}

func (o *StartRecordingParams) GetRequestType() string {
	return o.RequestType
}
func (o *StartRecordingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StartRecordingResponse represents the response body for the "StartRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartRecording.
*/
type StartRecordingResponse struct {
	api.Response
}

func (o *StartRecordingResponse) GetMessageID() string {
	return o.MessageID
}
func (o *StartRecordingResponse) GetStatus() string {
	return o.Status
}
func (o *StartRecordingResponse) GetError() string {
	return o.Error
}

/*
StopRecordingParams represents the params body for the "StopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopRecording.
*/
type StopRecordingParams struct {
	api.Params
}

func (o *StopRecordingParams) GetRequestType() string {
	return o.RequestType
}
func (o *StopRecordingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
StopRecordingResponse represents the response body for the "StopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopRecording.
*/
type StopRecordingResponse struct {
	api.Response
}

func (o *StopRecordingResponse) GetMessageID() string {
	return o.MessageID
}
func (o *StopRecordingResponse) GetStatus() string {
	return o.Status
}
func (o *StopRecordingResponse) GetError() string {
	return o.Error
}

/*
SetRecordingFolderParams represents the params body for the "SetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetRecordingFolder.
*/
type SetRecordingFolderParams struct {
	api.Params

	// Path of the recording folder.
	RecFolder string `json:"rec-folder"`
}

func (o *SetRecordingFolderParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetRecordingFolderParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetRecordingFolderResponse represents the response body for the "SetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetRecordingFolder.
*/
type SetRecordingFolderResponse struct {
	api.Response
}

func (o *SetRecordingFolderResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetRecordingFolderResponse) GetStatus() string {
	return o.Status
}
func (o *SetRecordingFolderResponse) GetError() string {
	return o.Error
}

/*
GetRecordingFolderParams represents the params body for the "GetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetRecordingFolder.
*/
type GetRecordingFolderParams struct {
	api.Params
}

func (o *GetRecordingFolderParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetRecordingFolderParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetRecordingFolderResponse represents the response body for the "GetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetRecordingFolder.
*/
type GetRecordingFolderResponse struct {
	api.Response

	// Path of the recording folder.
	RecFolder string `json:"rec-folder"`
}

func (o *GetRecordingFolderResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetRecordingFolderResponse) GetStatus() string {
	return o.Status
}
func (o *GetRecordingFolderResponse) GetError() string {
	return o.Error
}
