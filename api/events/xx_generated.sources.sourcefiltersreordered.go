// This file has been automatically generated. Don't edit it.

package events

/*
SourceFiltersReordered represents the event body for the "SourceFiltersReordered" event.
Since v4.6.0.
*/
type SourceFiltersReordered struct {
	EventBasic

	Filters []struct {
		// Filter visibility status
		Enabled bool `json:"enabled"`

		// Filter name
		Name string `json:"name"`

		// Filter type
		Type string `json:"type"`
	} `json:"filters"`

	// Source name
	SourceName string `json:"sourceName"`
}
