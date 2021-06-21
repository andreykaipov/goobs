// This file has been automatically generated. Don't edit it.

package events

/*
SourceFilterAdded represents the event body for the "SourceFilterAdded" event.
Since v4.6.0.
*/
type SourceFilterAdded struct {
	EventBasic

	// Filter name
	FilterName string `json:"filterName,omitempty"`

	// Filter settings
	FilterSettings map[string]interface{} `json:"filterSettings,omitempty"`

	// Filter type
	FilterType string `json:"filterType,omitempty"`

	// Source name
	SourceName string `json:"sourceName,omitempty"`
}
