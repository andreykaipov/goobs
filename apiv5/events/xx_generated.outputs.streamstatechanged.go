// This file has been automatically generated. Don't edit it.

package events

/*
StreamStateChanged represents the event body for the "StreamStateChanged" event.
Since v5.0.0.
*/
type StreamStateChanged struct {
	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`

	// The specific state of the output
	OutputState string `json:"outputState,omitempty"`
}
