// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
Represents the event body for the InputAudioTracksChanged event.

The audio tracks of an input have changed.
*/
type InputAudioTracksChanged struct {
	// Object of audio tracks along with their associated enable states
	InputAudioTracks *typedefs.InputAudioTracks `json:"inputAudioTracks,omitempty"`

	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid string `json:"inputUuid,omitempty"`
}
