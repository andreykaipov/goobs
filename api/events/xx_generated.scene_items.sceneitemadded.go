// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemAdded represents the event body for the "SceneItemAdded" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SceneItemAdded.
*/
type SceneItemAdded struct {
	EventBasic

	// Scene item ID
	ItemId int `json:"item-id"`

	// Name of the item added to the scene.
	ItemName string `json:"item-name"`

	// Name of the scene.
	SceneName string `json:"scene-name"`
}
