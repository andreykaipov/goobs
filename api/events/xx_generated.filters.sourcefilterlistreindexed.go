// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the event body for the SourceFilterListReindexed event.
type SourceFilterListReindexed struct {
	Filters []*typedefs.Filter `json:"filters,omitempty"`

	// Name of the source
	SourceName string `json:"sourceName,omitempty"`
}
