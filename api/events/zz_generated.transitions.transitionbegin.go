// This file has been automatically generated. Don't edit it.

package events

/*
TransitionBegin represents the event body for the "TransitionBegin" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#TransitionBegin.
*/
type TransitionBegin struct {
	EventBasic

	// Transition duration (in milliseconds).
	Duration int `json:"duration"`

	// Source scene of the transition
	FromScene string `json:"from-scene"`

	// Transition name.
	Name string `json:"name"`

	// Destination scene of the transition
	ToScene string `json:"to-scene"`
}
