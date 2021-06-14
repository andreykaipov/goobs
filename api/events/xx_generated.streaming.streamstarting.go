// This file has been automatically generated. Don't edit it.

package events

/*
StreamStarting represents the event body for the "StreamStarting" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#StreamStarting.
*/
type StreamStarting struct {
	EventBasic

	// Always false (retrocompatibility).
	PreviewOnly bool `json:"preview-only"`
}
