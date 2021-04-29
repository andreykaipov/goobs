// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemAdded represents the event body for the "SceneItemAdded" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SceneItemAdded.
*/
type SceneItemAdded struct {
	EventBasic

	// Name of the item added to the scene.
	ItemName string `json:"item-name"`

	// Name of the scene.
	SceneName string `json:"scene-name"`
}
