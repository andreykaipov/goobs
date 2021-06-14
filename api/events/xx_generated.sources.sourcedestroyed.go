// This file has been automatically generated. Don't edit it.

package events

/*
SourceDestroyed represents the event body for the "SourceDestroyed" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SourceDestroyed.
*/
type SourceDestroyed struct {
	EventBasic

	// Source kind.
	SourceKind string `json:"sourceKind"`

	// Source name
	SourceName string `json:"sourceName"`

	// Source type. Can be "input", "scene", "transition" or "filter".
	SourceType string `json:"sourceType"`
}
