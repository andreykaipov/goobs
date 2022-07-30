// This file has been automatically generated. Don't edit it.

package events

/*
InputShowStateChanged represents the event body for the "InputShowStateChanged" event.
Since v5.0.0.
*/
type InputShowStateChanged struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// Whether the input is showing
	VideoShowing bool `json:"videoShowing,omitempty"`
}
