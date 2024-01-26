// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the InputActiveStateChanged event.

An input's active state has changed.

When an input is active, it means it's being shown by the program feed.
*/
type InputActiveStateChanged struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid string `json:"inputUuid,omitempty"`

	// Whether the input is active
	VideoActive bool `json:"videoActive,omitempty"`
}
