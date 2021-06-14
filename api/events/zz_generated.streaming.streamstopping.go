// This file has been automatically generated. Don't edit it.

package events

/*
StreamStopping represents the event body for the "StreamStopping" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#StreamStopping.
*/
type StreamStopping struct {
	EventBasic

	// Always false (retrocompatibility).
	PreviewOnly bool `json:"preview-only"`
}
