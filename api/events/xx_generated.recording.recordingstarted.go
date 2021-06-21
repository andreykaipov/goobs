// This file has been automatically generated. Don't edit it.

package events

/*
RecordingStarted represents the event body for the "RecordingStarted" event.
Since v0.3.
*/
type RecordingStarted struct {
	EventBasic

	// Absolute path to the file of the current recording.
	RecordingFilename string `json:"recordingFilename,omitempty"`
}
