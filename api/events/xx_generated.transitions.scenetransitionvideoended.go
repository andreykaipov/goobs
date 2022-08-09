// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the SceneTransitionVideoEnded event.

A scene transition's video has completed fully.

Useful for stinger transitions to tell when the video *actually* ends.
`SceneTransitionEnded` only signifies the cut point, not the completion of transition playback.

Note: Appears to be called by every transition, regardless of relevance.
*/
type SceneTransitionVideoEnded struct {
	// Scene transition name
	TransitionName string `json:"transitionName,omitempty"`
}
