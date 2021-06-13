// This file has been automatically generated. Don't edit it.

package events

/*
SourceFilterAdded represents the event body for the "SourceFilterAdded" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#SourceFilterAdded.
*/
type SourceFilterAdded struct {
	EventBasic

	// Filter name
	FilterName string `json:"filterName"`

	// Filter settings
	FilterSettings map[string]interface{} `json:"filterSettings"`

	// Filter type
	FilterType string `json:"filterType"`

	// Source name
	SourceName string `json:"sourceName"`
}
