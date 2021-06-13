// This file has been automatically generated. Don't edit it.

package events

/*
SourceAudioSyncOffsetChanged represents the event body for the "SourceAudioSyncOffsetChanged" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#SourceAudioSyncOffsetChanged.
*/
type SourceAudioSyncOffsetChanged struct {
	EventBasic

	// Source name
	SourceName string `json:"sourceName"`

	// Audio sync offset of the source (in nanoseconds)
	SyncOffset int `json:"syncOffset"`
}
