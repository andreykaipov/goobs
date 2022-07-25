// This file has been automatically generated. Don't edit it.

package events

/*
SourceFilterEnableStateChanged represents the event body for the "SourceFilterEnableStateChanged" event.
Since v5.0.0.
*/
type SourceFilterEnableStateChanged struct {
	// Whether the filter is enabled
	FilterEnabled bool `json:"filterEnabled,omitempty"`

	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}
