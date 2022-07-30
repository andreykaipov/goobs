// This file has been automatically generated. Don't edit it.

package events

/*
InputActiveStateChanged represents the event body for the "InputActiveStateChanged" event.
Since v5.0.0.
*/
type InputActiveStateChanged struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// Whether the input is active
	VideoActive bool `json:"videoActive,omitempty"`
}
