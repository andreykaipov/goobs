// This file has been automatically generated. Don't edit it.

package events

/*
StreamStopping represents the event body for the "StreamStopping" event.
Since v0.3.
*/
type StreamStopping struct {
	EventBasic

	// Always false (retrocompatibility).
	PreviewOnly bool `json:"preview-only"`
}
