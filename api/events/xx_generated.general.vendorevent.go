// This file has been automatically generated. Don't edit it.

package events

/*
VendorEvent represents the event body for the "VendorEvent" event.
Since v5.0.0.
*/
type VendorEvent struct {
	// Vendor-provided event data. {} if event does not provide any data
	EventData interface{} `json:"eventData,omitempty"`

	// Vendor-provided event typedef
	EventType string `json:"eventType,omitempty"`

	// Name of the vendor emitting the event
	VendorName string `json:"vendorName,omitempty"`
}
