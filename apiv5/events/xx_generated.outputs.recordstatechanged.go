// This file has been automatically generated. Don't edit it.

package events

/*
RecordStateChanged represents the event body for the "RecordStateChanged" event.
Since v5.0.0.
*/
type RecordStateChanged struct {
	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`

	// The specific state of the output
	OutputState string `json:"outputState,omitempty"`
}
