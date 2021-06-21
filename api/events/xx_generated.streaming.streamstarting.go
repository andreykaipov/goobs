// This file has been automatically generated. Don't edit it.

package events

/*
StreamStarting represents the event body for the "StreamStarting" event.
Since v0.3.
*/
type StreamStarting struct {
	EventBasic

	// Always false (retrocompatibility).
	PreviewOnly bool `json:"preview-only"`
}
