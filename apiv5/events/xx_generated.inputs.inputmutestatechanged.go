// This file has been automatically generated. Don't edit it.

package events

/*
InputMuteStateChanged represents the event body for the "InputMuteStateChanged" event.
Since v5.0.0.
*/
type InputMuteStateChanged struct {
	// Whether the input is muted
	InputMuted bool `json:"inputMuted,omitempty"`

	// Name of the input
	InputName string `json:"inputName,omitempty"`
}
