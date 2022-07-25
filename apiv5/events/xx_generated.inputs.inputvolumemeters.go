// This file has been automatically generated. Don't edit it.

package events

/*
InputVolumeMeters represents the event body for the "InputVolumeMeters" event.
Since v5.0.0.
*/
type InputVolumeMeters struct {
	// Array of active inputs with their associated volume levels
	Inputs []interface{} `json:"inputs,omitempty"`
}
