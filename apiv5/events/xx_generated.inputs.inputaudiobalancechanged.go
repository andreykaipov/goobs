// This file has been automatically generated. Don't edit it.

package events

/*
InputAudioBalanceChanged represents the event body for the "InputAudioBalanceChanged" event.
Since v5.0.0.
*/
type InputAudioBalanceChanged struct {
	// New audio balance value of the input
	InputAudioBalance float64 `json:"inputAudioBalance,omitempty"`

	// Name of the affected input
	InputName string `json:"inputName,omitempty"`
}
