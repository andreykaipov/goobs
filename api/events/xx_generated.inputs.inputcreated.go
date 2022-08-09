// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the InputCreated event.

An input has been created.
*/
type InputCreated struct {
	// The default settings for the input
	DefaultInputSettings map[string]interface{} `json:"defaultInputSettings,omitempty"`

	// The kind of the input
	InputKind string `json:"inputKind,omitempty"`

	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// The settings configured to the input when it was created
	InputSettings map[string]interface{} `json:"inputSettings,omitempty"`

	// The unversioned kind of input (aka no `_v2` stuff)
	UnversionedInputKind string `json:"unversionedInputKind,omitempty"`
}
