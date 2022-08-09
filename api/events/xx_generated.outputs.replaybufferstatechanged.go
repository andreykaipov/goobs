// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the ReplayBufferStateChanged event.

The state of the replay buffer output has changed.
*/
type ReplayBufferStateChanged struct {
	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`

	// The specific state of the output
	OutputState string `json:"outputState,omitempty"`
}
