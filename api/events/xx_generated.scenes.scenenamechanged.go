// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the SceneNameChanged event.

The name of a scene has changed.
*/
type SceneNameChanged struct {
	// Old name of the scene
	OldSceneName string `json:"oldSceneName,omitempty"`

	// New name of the scene
	SceneName string `json:"sceneName,omitempty"`

	// UUID of the scene
	SceneUuid string `json:"sceneUuid,omitempty"`
}
