// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the InputShowStateChanged event.

An input's show state has changed.

When an input is showing, it means it's being shown by the preview or a dialog.
*/
type InputShowStateChanged struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid string `json:"inputUuid,omitempty"`

	// Whether the input is showing
	VideoShowing bool `json:"videoShowing,omitempty"`
}
