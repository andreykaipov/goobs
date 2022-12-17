// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the RecordStateChanged event.

The state of the record output has changed.
*/
type RecordStateChanged struct {
	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`

	// File name for the saved recording, if record stopped. `null` otherwise
	OutputPath string `json:"outputPath,omitempty"`

	// The specific state of the output
	OutputState string `json:"outputState,omitempty"`
}
