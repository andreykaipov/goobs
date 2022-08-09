// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the InputNameChanged event.

The name of an input has changed.
*/
type InputNameChanged struct {
	// New name of the input
	InputName string `json:"inputName,omitempty"`

	// Old name of the input
	OldInputName string `json:"oldInputName,omitempty"`
}
