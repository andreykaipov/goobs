// This file has been automatically generated. Don't edit it.

package events

// Represents the event body for the InputActiveStateChanged event.
type InputActiveStateChanged struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// Whether the input is active
	VideoActive bool `json:"videoActive,omitempty"`
}
