// This file has been automatically generated. Don't edit it.

package events

/*
InputCreated represents the event body for the "InputCreated" event.
Since v5.0.0.
*/
type InputCreated struct {
	// The default settings for the input
	DefaultInputSettings interface{} `json:"defaultInputSettings,omitempty"`

	// The kind of the input
	InputKind string `json:"inputKind,omitempty"`

	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// The settings configured to the input when it was created
	InputSettings interface{} `json:"inputSettings,omitempty"`

	// The unversioned kind of input (aka no `_v2` stuff)
	UnversionedInputKind string `json:"unversionedInputKind,omitempty"`
}
