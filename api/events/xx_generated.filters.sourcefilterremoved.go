// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the SourceFilterRemoved event.

A filter has been removed from a source.
*/
type SourceFilterRemoved struct {
	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// Name of the source the filter was on
	SourceName string `json:"sourceName,omitempty"`
}
