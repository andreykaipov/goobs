// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the CanvasRemoved event.

A canvas has been removed.
*/
type CanvasRemoved struct {
	// Name of the removed canvas
	CanvasName string `json:"canvasName,omitempty"`

	// UUID of the removed canvas
	CanvasUuid string `json:"canvasUuid,omitempty"`
}
