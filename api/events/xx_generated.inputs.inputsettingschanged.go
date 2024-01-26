// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the InputSettingsChanged event.

An input's settings have changed (been updated).

Note: On some inputs, changing values in the properties dialog will cause an immediate update. Pressing the "Cancel" button will revert the settings, resulting in another event being fired.
*/
type InputSettingsChanged struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// New settings object of the input
	InputSettings map[string]any `json:"inputSettings,omitempty"`

	// UUID of the input
	InputUuid string `json:"inputUuid,omitempty"`
}
