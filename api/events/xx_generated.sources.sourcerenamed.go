// This file has been automatically generated. Don't edit it.

package events

/*
SourceRenamed represents the event body for the "SourceRenamed" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SourceRenamed.
*/
type SourceRenamed struct {
	EventBasic

	// New source name
	NewName string `json:"newName"`

	// Previous source name
	PreviousName string `json:"previousName"`

	// Type of source (input, scene, filter, transition)
	SourceType string `json:"sourceType"`
}
