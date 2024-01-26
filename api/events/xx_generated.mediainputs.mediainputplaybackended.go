// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the MediaInputPlaybackEnded event.

A media input has finished playing.
*/
type MediaInputPlaybackEnded struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid string `json:"inputUuid,omitempty"`
}
