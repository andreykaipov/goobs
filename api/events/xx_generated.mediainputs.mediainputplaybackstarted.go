// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the MediaInputPlaybackStarted event.

A media input has started playing.
*/
type MediaInputPlaybackStarted struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid string `json:"inputUuid,omitempty"`
}
