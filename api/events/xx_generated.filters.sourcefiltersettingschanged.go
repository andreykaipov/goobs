// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the SourceFilterSettingsChanged event.

An source filter's settings have changed (been updated).
*/
type SourceFilterSettingsChanged struct {
	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// New settings object of the filter
	FilterSettings map[string]any `json:"filterSettings,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}
