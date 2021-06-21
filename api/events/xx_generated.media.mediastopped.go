// This file has been automatically generated. Don't edit it.

package events

/*
MediaStopped represents the event body for the "MediaStopped" event.
Since v4.9.0.
*/
type MediaStopped struct {
	EventBasic

	// The ID type of the source (Eg. `vlc_source` or `ffmpeg_source`)
	SourceKind string `json:"sourceKind,omitempty"`

	// Source name
	SourceName string `json:"sourceName,omitempty"`
}
