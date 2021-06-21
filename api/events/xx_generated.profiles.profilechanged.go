// This file has been automatically generated. Don't edit it.

package events

/*
ProfileChanged represents the event body for the "ProfileChanged" event.
Since v4.0.0.
*/
type ProfileChanged struct {
	EventBasic

	// Name of the new current profile.
	Profile string `json:"profile,omitempty"`
}
