// This file has been automatically generated. Don't edit it.

package events

/*
SourceAudioActivated represents the event body for the "SourceAudioActivated" event.
Since v4.9.0.
*/
type SourceAudioActivated struct {
	EventBasic

	// Source name
	SourceName string `json:"sourceName,omitempty"`
}
