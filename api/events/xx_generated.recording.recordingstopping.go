// This file has been automatically generated. Don't edit it.

package events

/*
RecordingStopping represents the event body for the "RecordingStopping" event.
Since v0.3.
*/
type RecordingStopping struct {
	EventBasic

	// Absolute path to the file of the current recording.
	RecordingFilename string `json:"recordingFilename,omitempty"`
}
