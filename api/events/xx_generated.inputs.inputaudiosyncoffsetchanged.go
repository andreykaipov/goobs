// This file has been automatically generated. Don't edit it.

package events

/*
InputAudioSyncOffsetChanged represents the event body for the "InputAudioSyncOffsetChanged" event.
Since v5.0.0.
*/
type InputAudioSyncOffsetChanged struct {
	// New sync offset in milliseconds
	InputAudioSyncOffset float64 `json:"inputAudioSyncOffset,omitempty"`

	// Name of the input
	InputName string `json:"inputName,omitempty"`
}
