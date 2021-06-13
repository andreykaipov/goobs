// This file has been automatically generated. Don't edit it.

package events

/*
SourceFiltersReordered represents the event body for the "SourceFiltersReordered" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#SourceFiltersReordered.
*/
type SourceFiltersReordered struct {
	EventBasic

	Filters []struct {
		// Filter name
		Name string `json:"name"`

		// Filter type
		Type string `json:"type"`
	} `json:"filters"`

	// Source name
	SourceName string `json:"sourceName"`
}
