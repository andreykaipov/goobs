// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the SourceFilterEnableStateChanged event.

A source filter's enable state has changed.
*/
type SourceFilterEnableStateChanged struct {
	// Whether the filter is enabled
	FilterEnabled bool `json:"filterEnabled,omitempty"`

	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}
