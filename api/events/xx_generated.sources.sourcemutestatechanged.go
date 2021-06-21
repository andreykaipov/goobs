// This file has been automatically generated. Don't edit it.

package events

/*
SourceMuteStateChanged represents the event body for the "SourceMuteStateChanged" event.
Since v4.6.0.
*/
type SourceMuteStateChanged struct {
	EventBasic

	// Mute status of the source
	Muted bool `json:"muted"`

	// Source name
	SourceName string `json:"sourceName,omitempty"`
}
