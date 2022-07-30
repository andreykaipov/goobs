// This file has been automatically generated. Don't edit it.

package events

/*
SceneRemoved represents the event body for the "SceneRemoved" event.
Since v5.0.0.
*/
type SceneRemoved struct {
	// Whether the scene was a group
	IsGroup bool `json:"isGroup,omitempty"`

	// Name of the removed scene
	SceneName string `json:"sceneName,omitempty"`
}
