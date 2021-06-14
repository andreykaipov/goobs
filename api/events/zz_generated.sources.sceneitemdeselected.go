// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemDeselected represents the event body for the "SceneItemDeselected" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#SceneItemDeselected.
*/
type SceneItemDeselected struct {
	EventBasic

	// Name of the item in the scene.
	ItemId int `json:"item-id"`

	// Name of the item in the scene.
	ItemName string `json:"item-name"`

	// Name of the scene.
	SceneName string `json:"scene-name"`
}
