// This file has been automatically generated. Don't edit it.

package events

/*
SceneCreated represents the event body for the "SceneCreated" event.
Since v5.0.0.
*/
type SceneCreated struct {
	// Whether the new scene is a group
	IsGroup bool `json:"isGroup,omitempty"`

	// Name of the new scene
	SceneName string `json:"sceneName,omitempty"`
}
