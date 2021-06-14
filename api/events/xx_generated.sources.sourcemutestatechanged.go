// This file has been automatically generated. Don't edit it.

package events

/*
SourceMuteStateChanged represents the event body for the "SourceMuteStateChanged" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SourceMuteStateChanged.
*/
type SourceMuteStateChanged struct {
	EventBasic

	// Mute status of the source
	Muted bool `json:"muted"`

	// Source name
	SourceName string `json:"sourceName"`
}
