// This file has been automatically generated. Don't edit it.

package events

/*
RecordingStopped represents the event body for the "RecordingStopped" event.
Since v0.3.
*/
type RecordingStopped struct {
	EventBasic

	// Absolute path to the file of the current recording.
	RecordingFilename string `json:"recordingFilename,omitempty"`
}
