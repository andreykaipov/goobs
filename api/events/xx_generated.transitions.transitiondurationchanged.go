// This file has been automatically generated. Don't edit it.

package events

/*
TransitionDurationChanged represents the event body for the "TransitionDurationChanged" event.
Since v4.0.0.
*/
type TransitionDurationChanged struct {
	EventBasic

	// New transition duration.
	NewDuration int `json:"new-duration,omitempty"`
}
