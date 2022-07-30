// This file has been automatically generated. Don't edit it.

package events

/*
InputAudioTracksChanged represents the event body for the "InputAudioTracksChanged" event.
Since v5.0.0.
*/
type InputAudioTracksChanged struct {
	// Object of audio tracks along with their associated enable states
	InputAudioTracks interface{} `json:"inputAudioTracks,omitempty"`

	// Name of the input
	InputName string `json:"inputName,omitempty"`
}
