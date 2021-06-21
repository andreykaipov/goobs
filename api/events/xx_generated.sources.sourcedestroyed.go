// This file has been automatically generated. Don't edit it.

package events

/*
SourceDestroyed represents the event body for the "SourceDestroyed" event.
Since v4.6.0.
*/
type SourceDestroyed struct {
	EventBasic

	// Source kind.
	SourceKind string `json:"sourceKind,omitempty"`

	// Source name
	SourceName string `json:"sourceName,omitempty"`

	// Source type. Can be "input", "scene", "transition" or "filter".
	SourceType string `json:"sourceType,omitempty"`
}
