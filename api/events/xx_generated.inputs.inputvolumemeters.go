// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
Represents the event body for the InputVolumeMeters event.

A high-volume event providing volume levels of all active inputs every 50 milliseconds.
*/
type InputVolumeMeters struct {
	Inputs []*typedefs.Input `json:"inputs,omitempty"`
}
