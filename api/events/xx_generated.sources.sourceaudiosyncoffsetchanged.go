// This file has been automatically generated. Don't edit it.

package events

/*
SourceAudioSyncOffsetChanged represents the event body for the "SourceAudioSyncOffsetChanged" event.
Since v4.6.0.
*/
type SourceAudioSyncOffsetChanged struct {
	EventBasic

	// Source name
	SourceName string `json:"sourceName,omitempty"`

	// Audio sync offset of the source (in nanoseconds)
	SyncOffset int `json:"syncOffset,omitempty"`
}
