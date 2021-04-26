// This file has been automatically generated. Don't edit it.

package recording

// StartStopRecordingResponse contains the request body for the [StartStopRecording](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopRecording) request.
type StartStopRecordingResponse struct{}

// StartStopRecordingParams contains the request body for the [StartStopRecording](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopRecording) request.
type StartStopRecordingParams struct{}

// StartRecordingParams contains the request body for the [StartRecording](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartRecording) request.
type StartRecordingParams struct{}

// StartRecordingResponse contains the request body for the [StartRecording](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartRecording) request.
type StartRecordingResponse struct{}

// StopRecordingParams contains the request body for the [StopRecording](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopRecording) request.
type StopRecordingParams struct{}

// StopRecordingResponse contains the request body for the [StopRecording](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StopRecording) request.
type StopRecordingResponse struct{}

// SetRecordingFolderParams contains the request body for the [SetRecordingFolder](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetRecordingFolder) request.
type SetRecordingFolderParams struct {
	// Path of the recording folder.
	RecFolder string `json:"rec-folder"`
}

// SetRecordingFolderResponse contains the request body for the [SetRecordingFolder](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetRecordingFolder) request.
type SetRecordingFolderResponse struct{}

// GetRecordingFolderParams contains the request body for the [GetRecordingFolder](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetRecordingFolder) request.
type GetRecordingFolderParams struct{}

// GetRecordingFolderResponse contains the request body for the [GetRecordingFolder](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetRecordingFolder) request.
type GetRecordingFolderResponse struct {
	// Path of the recording folder.
	RecFolder string `json:"rec-folder"`
}
