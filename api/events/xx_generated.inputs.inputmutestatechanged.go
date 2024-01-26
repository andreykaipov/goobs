// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the InputMuteStateChanged event.

An input's mute state has changed.
*/
type InputMuteStateChanged struct {
	// Whether the input is muted
	InputMuted bool `json:"inputMuted,omitempty"`

	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid string `json:"inputUuid,omitempty"`
}
