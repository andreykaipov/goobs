// This file has been automatically generated. Don't edit it.

package events

/*
TransitionEnd represents the event body for the "TransitionEnd" event.
Since v4.8.0.
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
