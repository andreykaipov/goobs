// This file has been automatically generated. Don't edit it.

package events

/*
SourceAudioDeactivated represents the event body for the "SourceAudioDeactivated" event.
Since v4.9.0.
*/
type SourceAudioDeactivated struct {
	EventBasic

	// Source name
	SourceName string `json:"sourceName,omitempty"`
}
