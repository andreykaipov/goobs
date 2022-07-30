// This file has been automatically generated. Don't edit it.

package events

/*
VirtualcamStateChanged represents the event body for the "VirtualcamStateChanged" event.
Since v5.0.0.
*/
type VirtualcamStateChanged struct {
	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`

	// The specific state of the output
	OutputState string `json:"outputState,omitempty"`
}
