// This file has been automatically generated. Don't edit it.

package events

/*
BroadcastCustomMessage represents the event body for the "BroadcastCustomMessage" event.
Since v4.7.0.
*/
type BroadcastCustomMessage struct {
	EventBasic

	// User-defined data
	Data map[string]interface{} `json:"data,omitempty"`

	// Identifier provided by the sender
	Realm string `json:"realm,omitempty"`
}
