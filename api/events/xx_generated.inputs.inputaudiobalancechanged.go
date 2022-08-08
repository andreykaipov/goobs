// This file has been automatically generated. Don't edit it.

package events

// Represents the event body for the InputAudioBalanceChanged event.
type InputAudioBalanceChanged struct {
	// New audio balance value of the input
	InputAudioBalance float64 `json:"inputAudioBalance,omitempty"`

	// Name of the affected input
	InputName string `json:"inputName,omitempty"`
}
