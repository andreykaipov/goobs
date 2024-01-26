// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the SceneTransitionStarted event.

A scene transition has started.
*/
type SceneTransitionStarted struct {
	// Scene transition name
	TransitionName string `json:"transitionName,omitempty"`

	// Scene transition UUID
	TransitionUuid string `json:"transitionUuid,omitempty"`
}
