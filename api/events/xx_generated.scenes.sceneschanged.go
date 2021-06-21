// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
ScenesChanged represents the event body for the "ScenesChanged" event.
Since v0.3.
*/
type ScenesChanged struct {
	EventBasic

	// Scenes list.
	Scenes []typedefs.Scene `json:"scenes,omitempty"`
}
