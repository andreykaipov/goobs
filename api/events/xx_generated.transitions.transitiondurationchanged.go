// This file has been automatically generated. Don't edit it.

package events

/*
TransitionDurationChanged represents the event body for the "TransitionDurationChanged" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#TransitionDurationChanged.
*/
type TransitionDurationChanged struct {
	EventBasic

	// New transition duration.
	NewDuration int `json:"new-duration"`
}
