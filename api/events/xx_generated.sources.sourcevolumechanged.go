// This file has been automatically generated. Don't edit it.

package events

/*
SourceVolumeChanged represents the event body for the "SourceVolumeChanged" event.
Since v4.6.0.
*/
type SourceVolumeChanged struct {
	EventBasic

	// Source name
	SourceName string `json:"sourceName,omitempty"`

	// Source volume
	Volume float64 `json:"volume,omitempty"`

	// Source volume in Decibel
	VolumeDb float64 `json:"volumeDb,omitempty"`
}
