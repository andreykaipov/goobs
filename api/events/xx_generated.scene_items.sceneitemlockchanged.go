// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemLockChanged represents the event body for the "SceneItemLockChanged" event.
Since v4.8.0.
*/
type SceneItemLockChanged struct {
	EventBasic

	// Scene item ID
	ItemId int `json:"item-id,omitempty"`

	// New locked state of the item.
	ItemLocked bool `json:"item-locked"`

	// Name of the item in the scene.
	ItemName string `json:"item-name,omitempty"`

	// Name of the scene.
	SceneName string `json:"scene-name,omitempty"`
}
