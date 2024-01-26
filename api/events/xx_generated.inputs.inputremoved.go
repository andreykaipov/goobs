// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the InputRemoved event.

An input has been removed.
*/
type InputRemoved struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid string `json:"inputUuid,omitempty"`
}
