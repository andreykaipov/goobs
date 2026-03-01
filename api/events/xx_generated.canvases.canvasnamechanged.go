// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the CanvasNameChanged event.

The name of a canvas has changed.
*/
type CanvasNameChanged struct {
	// New name of the canvas
	CanvasName string `json:"canvasName,omitempty"`

	// UUID of the canvas
	CanvasUuid string `json:"canvasUuid,omitempty"`

	// Old name of the canvas
	OldCanvasName string `json:"oldCanvasName,omitempty"`
}
