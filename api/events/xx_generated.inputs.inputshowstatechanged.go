// This file has been automatically generated. Don't edit it.

package events

// Represents the event body for the InputShowStateChanged event.
type InputShowStateChanged struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// Whether the input is showing
	VideoShowing bool `json:"videoShowing,omitempty"`
}
