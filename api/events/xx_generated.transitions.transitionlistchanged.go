// This file has been automatically generated. Don't edit it.

package events

/*
TransitionListChanged represents the event body for the "TransitionListChanged" event.
Since v4.0.0.
*/
type TransitionListChanged struct {
	EventBasic

	Transitions []*Transition `json:"transitions,omitempty"`
}

type Transition struct {
	// Transition name.
	Name string `json:"name,omitempty"`
}
