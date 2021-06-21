// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
SourceFiltersReordered represents the event body for the "SourceFiltersReordered" event.
Since v4.6.0.
*/
type SourceFiltersReordered struct {
	EventBasic

	// The filters for this object.
	Filters []*typedefs.Filter `json:"filters,omitempty"`

	// Source name
	SourceName string `json:"sourceName,omitempty"`
}
