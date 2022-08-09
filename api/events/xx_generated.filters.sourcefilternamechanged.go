// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the SourceFilterNameChanged event.

The name of a source filter has changed.
*/
type SourceFilterNameChanged struct {
	// New name of the filter
	FilterName string `json:"filterName,omitempty"`

	// Old name of the filter
	OldFilterName string `json:"oldFilterName,omitempty"`

	// The source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}
