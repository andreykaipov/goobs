// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the CurrentSceneTransitionChanged event.

The current scene transition has changed.
*/
type CurrentSceneTransitionChanged struct {
	// Name of the new transition
	TransitionName string `json:"transitionName,omitempty"`

	// UUID of the new transition
	TransitionUuid string `json:"transitionUuid,omitempty"`
}
