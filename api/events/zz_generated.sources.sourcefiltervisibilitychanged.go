// This file has been automatically generated. Don't edit it.

package events

/*
SourceFilterVisibilityChanged represents the event body for the "SourceFilterVisibilityChanged" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#SourceFilterVisibilityChanged.
*/
type SourceFilterVisibilityChanged struct {
	EventBasic

	// New filter state
	FilterEnabled bool `json:"filterEnabled"`

	// Filter name
	FilterName string `json:"filterName"`

	// Source name
	SourceName string `json:"sourceName"`
}
