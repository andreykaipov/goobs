// This file has been automatically generated. Don't edit it.

package events

/*
MediaPrevious represents the event body for the "MediaPrevious" event.
Since v4.9.0.
*/
type MediaPrevious struct {
	EventBasic

	// The ID type of the source (Eg. `vlc_source` or `ffmpeg_source`)
	SourceKind string `json:"sourceKind,omitempty"`

	// Source name
	SourceName string `json:"sourceName,omitempty"`
}
