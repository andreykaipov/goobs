// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the InputAudioSyncOffsetChanged event.

The sync offset of an input has changed.
*/
type InputAudioSyncOffsetChanged struct {
	// New sync offset in milliseconds
	InputAudioSyncOffset float64 `json:"inputAudioSyncOffset,omitempty"`

	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid string `json:"inputUuid,omitempty"`
}
