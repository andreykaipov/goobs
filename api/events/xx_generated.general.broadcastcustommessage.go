// This file has been automatically generated. Don't edit it.

package events

/*
BroadcastCustomMessage represents the event body for the "BroadcastCustomMessage" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#BroadcastCustomMessage.
*/
type BroadcastCustomMessage struct {
	EventBasic

	// User-defined data
	Data map[string]interface{} `json:"data"`

	// Identifier provided by the sender
	Realm string `json:"realm"`
}
