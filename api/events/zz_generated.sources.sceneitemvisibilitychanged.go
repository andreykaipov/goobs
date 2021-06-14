// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemVisibilityChanged represents the event body for the "SceneItemVisibilityChanged" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#SceneItemVisibilityChanged.
*/
type SceneItemVisibilityChanged struct {
	EventBasic

	// Scene item ID
	ItemId int `json:"item-id"`

	// Name of the item in the scene.
	ItemName string `json:"item-name"`

	// New visibility state of the item.
	ItemVisible bool `json:"item-visible"`

	// Name of the scene.
	SceneName string `json:"scene-name"`
}
