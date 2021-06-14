// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemLockChanged represents the event body for the "SceneItemLockChanged" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SceneItemLockChanged.
*/
type SceneItemLockChanged struct {
	EventBasic

	// Scene item ID
	ItemId int `json:"item-id"`

	// New locked state of the item.
	ItemLocked bool `json:"item-locked"`

	// Name of the item in the scene.
	ItemName string `json:"item-name"`

	// Name of the scene.
	SceneName string `json:"scene-name"`
}
