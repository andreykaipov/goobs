// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the SceneCreated event.

A new scene has been created.
*/
type SceneCreated struct {
	// Whether the new scene is a group
	IsGroup bool `json:"isGroup,omitempty"`

	// Name of the new scene
	SceneName string `json:"sceneName,omitempty"`
}
