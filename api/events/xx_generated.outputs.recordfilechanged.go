// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the RecordFileChanged event.

The record output has started writing to a new file. For example, when a file split happens.
*/
type RecordFileChanged struct {
	// File name that the output has begun writing to
	NewOutputPath string `json:"newOutputPath,omitempty"`
}
