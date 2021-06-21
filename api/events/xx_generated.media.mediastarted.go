// This file has been automatically generated. Don't edit it.

package events

/*
MediaStarted represents the event body for the "MediaStarted" event.
Since v4.9.0.
*/
type MediaStarted struct {
	EventBasic

	// The ID type of the source (Eg. `vlc_source` or `ffmpeg_source`)
	SourceKind string `json:"sourceKind,omitempty"`

	// Source name
	SourceName string `json:"sourceName,omitempty"`
}
