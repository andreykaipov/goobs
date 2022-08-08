// This file has been automatically generated. Don't edit it.

package events

// Represents the event body for the VendorEvent event.
type VendorEvent struct {
	// Vendor-provided event data. {} if event does not provide any data
	EventData interface{} `json:"eventData,omitempty"`

	// Vendor-provided event typedef
	EventType string `json:"eventType,omitempty"`

	// Name of the vendor emitting the event
	VendorName string `json:"vendorName,omitempty"`
}
