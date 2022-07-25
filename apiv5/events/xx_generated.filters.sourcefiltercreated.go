// This file has been automatically generated. Don't edit it.

package events

/*
SourceFilterCreated represents the event body for the "SourceFilterCreated" event.
Since v5.0.0.
*/
type SourceFilterCreated struct {
	// The default settings for the filter
	DefaultFilterSettings interface{} `json:"defaultFilterSettings,omitempty"`

	// Index position of the filter
	FilterIndex float64 `json:"filterIndex,omitempty"`

	// The kind of the filter
	FilterKind string `json:"filterKind,omitempty"`

	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// The settings configured to the filter when it was created
	FilterSettings interface{} `json:"filterSettings,omitempty"`

	// Name of the source the filter was added to
	SourceName string `json:"sourceName,omitempty"`
}
