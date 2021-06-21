// This file has been automatically generated. Don't edit it.

package events

/*
SourceFilterVisibilityChanged represents the event body for the "SourceFilterVisibilityChanged" event.
Since v4.7.0.
*/
type SourceFilterVisibilityChanged struct {
	EventBasic

	// New filter state
	FilterEnabled bool `json:"filterEnabled"`

	// Filter name
	FilterName string `json:"filterName,omitempty"`

	// Source name
	SourceName string `json:"sourceName,omitempty"`
}
