// This file has been automatically generated. Don't edit it.

package events

/*
SourceVolumeChanged represents the event body for the "SourceVolumeChanged" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#SourceVolumeChanged.
*/
type SourceVolumeChanged struct {
	EventBasic

	// Source name
	SourceName string `json:"sourceName"`

	// Source volume
	Volume float64 `json:"volume"`
}
