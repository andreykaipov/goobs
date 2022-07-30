// This file has been automatically generated. Don't edit it.

package events

/*
SourceFilterRemoved represents the event body for the "SourceFilterRemoved" event.
Since v5.0.0.
*/
type SourceFilterRemoved struct {
	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// Name of the source the filter was on
	SourceName string `json:"sourceName,omitempty"`
}
