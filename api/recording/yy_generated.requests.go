// This file has been automatically generated. Don't edit it.

package recording

/*
StartStopRecordingParams represents the params body for the "StartStopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopRecording.
*/
type StartStopRecordingParams struct{}

/*
StartStopRecordingResponse represents the response body for the "StartStopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopRecording.
*/
type StartStopRecordingResponse struct{}

/*
StartRecordingParams represents the params body for the "StartRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartRecording.
*/
type StartRecordingParams struct{}

/*
StartRecordingResponse represents the response body for the "StartRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartRecording.
*/
type StartRecordingResponse struct{}

/*
StopRecordingParams represents the params body for the "StopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopRecording.
*/
type StopRecordingParams struct{}

/*
StopRecordingResponse represents the response body for the "StopRecording" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopRecording.
*/
type StopRecordingResponse struct{}

/*
SetRecordingFolderParams represents the params body for the "SetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetRecordingFolder.
*/
type SetRecordingFolderParams struct {
	// Path of the recording folder.
	RecFolder string `json:"rec-folder"`
}

/*
SetRecordingFolderResponse represents the response body for the "SetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetRecordingFolder.
*/
type SetRecordingFolderResponse struct{}

/*
GetRecordingFolderParams represents the params body for the "GetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetRecordingFolder.
*/
type GetRecordingFolderParams struct{}

/*
GetRecordingFolderResponse represents the response body for the "GetRecordingFolder" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetRecordingFolder.
*/
type GetRecordingFolderResponse struct {
	// Path of the recording folder.
	RecFolder string `json:"rec-folder"`
}
