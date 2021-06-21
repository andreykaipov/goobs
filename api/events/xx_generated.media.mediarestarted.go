// This file has been automatically generated. Don't edit it.

package events

/*
MediaRestarted represents the event body for the "MediaRestarted" event.
Since v4.9.0.
*/
type MediaRestarted struct {
	EventBasic

	// The ID type of the source (Eg. `vlc_source` or `ffmpeg_source`)
	SourceKind string `json:"sourceKind,omitempty"`

	// Source name
	SourceName string `json:"sourceName,omitempty"`
}
