// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the InputVolumeChanged event.

An input's volume level has changed.
*/
type InputVolumeChanged struct {
	// Name of the input
	InputName string `json:"inputName,omitempty"`

	// UUID of the input
	InputUuid string `json:"inputUuid,omitempty"`

	// New volume level in dB
	InputVolumeDb float64 `json:"inputVolumeDb,omitempty"`

	// New volume level multiplier
	InputVolumeMul float64 `json:"inputVolumeMul,omitempty"`
}
