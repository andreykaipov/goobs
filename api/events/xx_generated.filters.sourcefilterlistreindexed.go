// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
SourceFilterListReindexed represents the event body for the "SourceFilterListReindexed" event.
Since v5.0.0.
*/
type SourceFilterListReindexed struct {
	Filters []*typedefs.Filter `json:"filters,omitempty"`

	// Name of the source
	SourceName string `json:"sourceName,omitempty"`
}
