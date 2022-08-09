// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the SceneItemLockStateChanged event.

A scene item's lock state has changed.
*/
type SceneItemLockStateChanged struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Whether the scene item is locked
	SceneItemLocked bool `json:"sceneItemLocked,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}
