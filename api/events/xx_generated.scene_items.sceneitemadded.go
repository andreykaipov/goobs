// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemAdded represents the event body for the "SceneItemAdded" event.
Since v4.0.0.
*/
type SceneItemAdded struct {
	EventBasic

	// Scene item ID
	ItemId int `json:"item-id,omitempty"`

	// Name of the item added to the scene.
	ItemName string `json:"item-name,omitempty"`

	// Name of the scene.
	SceneName string `json:"scene-name,omitempty"`
}
