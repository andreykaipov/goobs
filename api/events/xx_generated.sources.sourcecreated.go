// This file has been automatically generated. Don't edit it.

package events

/*
SourceCreated represents the event body for the "SourceCreated" event.
Since v4.6.0.
*/
type SourceCreated struct {
	EventBasic

	// Source kind.
	SourceKind string `json:"sourceKind"`

	// Source name
	SourceName string `json:"sourceName"`

	// Source settings
	SourceSettings map[string]interface{} `json:"sourceSettings"`

	// Source type. Can be "input", "scene", "transition" or "filter".
	SourceType string `json:"sourceType"`
}
