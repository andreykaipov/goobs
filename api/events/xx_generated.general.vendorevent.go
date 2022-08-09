// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the VendorEvent event.

An event has been emitted from a vendor.

A vendor is a unique name registered by a third-party plugin or script, which allows for custom requests and events to be added to obs-websocket.
If a plugin or script implements vendor requests or events, documentation is expected to be provided with them.
*/
type VendorEvent struct {
	// Vendor-provided event data. {} if event does not provide any data
	EventData interface{} `json:"eventData,omitempty"`

	// Vendor-provided event typedef
	EventType string `json:"eventType,omitempty"`

	// Name of the vendor emitting the event
	VendorName string `json:"vendorName,omitempty"`
}
