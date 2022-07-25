// This file has been automatically generated. Don't edit it.

package events

/*
InputNameChanged represents the event body for the "InputNameChanged" event.
Since v5.0.0.
*/
type InputNameChanged struct {
	// New name of the input
	InputName string `json:"inputName,omitempty"`

	// Old name of the input
	OldInputName string `json:"oldInputName,omitempty"`
}
