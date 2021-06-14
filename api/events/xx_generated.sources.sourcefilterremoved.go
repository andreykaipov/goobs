// This file has been automatically generated. Don't edit it.

package events

/*
SourceFilterRemoved represents the event body for the "SourceFilterRemoved" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#SourceFilterRemoved.
*/
type SourceFilterRemoved struct {
	EventBasic

	// Filter name
	FilterName string `json:"filterName"`

	// Filter type
	FilterType string `json:"filterType"`

	// Source name
	SourceName string `json:"sourceName"`
}
