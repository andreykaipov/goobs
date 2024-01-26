// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the SceneRemoved event.

A scene has been removed.
*/
type SceneRemoved struct {
	// Whether the scene was a group
	IsGroup bool `json:"isGroup,omitempty"`

	// Name of the removed scene
	SceneName string `json:"sceneName,omitempty"`

	// UUID of the removed scene
	SceneUuid string `json:"sceneUuid,omitempty"`
}
