// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the CanvasCreated event.

A new canvas has been created.
*/
type CanvasCreated struct {
	// Name of the new canvas
	CanvasName string `json:"canvasName,omitempty"`

	// UUID of the new canvas
	CanvasUuid string `json:"canvasUuid,omitempty"`
}
