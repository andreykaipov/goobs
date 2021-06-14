// This file has been automatically generated. Don't edit it.

package events

/*
SwitchTransition represents the event body for the "SwitchTransition" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SwitchTransition.
*/
type SwitchTransition struct {
	EventBasic

	// The name of the new active transition.
	TransitionName string `json:"transition-name"`
}
