// This file has been automatically generated. Don't edit it.

package events

/*
SwitchTransition represents the event body for the "SwitchTransition" event.
Since v4.0.0.
*/
type SwitchTransition struct {
	EventBasic

	// The name of the new active transition.
	TransitionName string `json:"transition-name,omitempty"`
}
