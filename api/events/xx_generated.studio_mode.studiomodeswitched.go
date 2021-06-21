// This file has been automatically generated. Don't edit it.

package events

/*
StudioModeSwitched represents the event body for the "StudioModeSwitched" event.
Since v4.1.0.
*/
type StudioModeSwitched struct {
	EventBasic

	// The new enabled state of Studio Mode.
	NewState bool `json:"new-state"`
}
