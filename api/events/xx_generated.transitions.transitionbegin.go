// This file has been automatically generated. Don't edit it.

package events

/*
TransitionBegin represents the event body for the "TransitionBegin" event.
Since v4.0.0.
*/
type TransitionBegin struct {
	EventBasic

	// Transition duration (in milliseconds). Will be -1 for any transition with a fixed duration, such as a Stinger,
	// due to limitations of the OBS API.
	Duration int `json:"duration,omitempty"`

	// Source scene of the transition
	FromScene string `json:"from-scene,omitempty"`

	// Transition name.
	Name string `json:"name,omitempty"`

	// Destination scene of the transition
	ToScene string `json:"to-scene,omitempty"`

	// Transition type.
	Type string `json:"type,omitempty"`
}
