// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the VirtualcamStateChanged event.

The state of the virtualcam output has changed.
*/
type VirtualcamStateChanged struct {
	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`

	// The specific state of the output
	OutputState string `json:"outputState,omitempty"`
}
