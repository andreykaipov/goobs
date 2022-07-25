// This file has been automatically generated. Don't edit it.

package events

/*
SceneNameChanged represents the event body for the "SceneNameChanged" event.
Since v5.0.0.
*/
type SceneNameChanged struct {
	// Old name of the scene
	OldSceneName string `json:"oldSceneName,omitempty"`

	// New name of the scene
	SceneName string `json:"sceneName,omitempty"`
}
