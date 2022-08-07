// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
InputAudioTracksChanged represents the event body for the "InputAudioTracksChanged" event.
Since v5.0.0.
*/
type InputAudioTracksChanged struct {
	InputAudioTracks *typedefs.InputAudioTracks `json:"inputAudioTracks,omitempty"`

	// Name of the input
	InputName string `json:"inputName,omitempty"`
}
