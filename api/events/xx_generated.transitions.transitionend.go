// This file has been automatically generated. Don't edit it.

package events

/*
TransitionEnd represents the event body for the "TransitionEnd" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#TransitionEnd.
*/
type TransitionEnd struct {
	EventBasic

	// Transition duration (in milliseconds).
	Duration int `json:"duration"`

	// Transition name.
	Name string `json:"name"`

	// Destination scene of the transition
	ToScene string `json:"to-scene"`

	// Transition type.
	Type string `json:"type"`
}
