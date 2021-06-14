// This file has been automatically generated. Don't edit it.

package events

/*
TransitionBegin represents the event body for the "TransitionBegin" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#TransitionBegin.
*/
type TransitionBegin struct {
	EventBasic

	// Transition duration (in milliseconds). Will be -1 for any transition with a fixed duration,
	// such as a Stinger, due to limitations of the OBS API.
	Duration int `json:"duration"`

	// Source scene of the transition
	FromScene string `json:"from-scene"`

	// Transition name.
	Name string `json:"name"`

	// Destination scene of the transition
	ToScene string `json:"to-scene"`

	// Transition type.
	Type string `json:"type"`
}
