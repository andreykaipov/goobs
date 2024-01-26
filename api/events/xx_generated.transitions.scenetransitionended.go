// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the SceneTransitionEnded event.

A scene transition has completed fully.

Note: Does not appear to trigger when the transition is interrupted by the user.
*/
type SceneTransitionEnded struct {
	// Scene transition name
	TransitionName string `json:"transitionName,omitempty"`

	// Scene transition UUID
	TransitionUuid string `json:"transitionUuid,omitempty"`
}
