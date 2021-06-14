// This file has been automatically generated. Don't edit it.

package events

/*
StudioModeSwitched represents the event body for the "StudioModeSwitched" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#StudioModeSwitched.
*/
type StudioModeSwitched struct {
	EventBasic

	// The new enabled state of Studio Mode.
	NewState bool `json:"new-state"`
}
