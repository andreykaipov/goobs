// This file has been automatically generated. Don't edit it.

package events

/*
SourceOrderChanged represents the event body for the "SourceOrderChanged" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SourceOrderChanged.
*/
type SourceOrderChanged struct {
	EventBasic

	// Name of the scene where items have been reordered.
	SceneName string `json:"scene-name"`
}
