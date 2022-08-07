// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
InputVolumeMeters represents the event body for the "InputVolumeMeters" event.
Since v5.0.0.
*/
type InputVolumeMeters struct {
	Inputs []*typedefs.Input `json:"inputs,omitempty"`
}
