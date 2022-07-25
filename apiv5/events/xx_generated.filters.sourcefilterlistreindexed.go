// This file has been automatically generated. Don't edit it.

package events

/*
SourceFilterListReindexed represents the event body for the "SourceFilterListReindexed" event.
Since v5.0.0.
*/
type SourceFilterListReindexed struct {
	// Array of filter objects
	Filters []interface{} `json:"filters,omitempty"`

	// Name of the source
	SourceName string `json:"sourceName,omitempty"`
}
