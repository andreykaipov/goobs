// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemVisibilityChanged represents the event body for the "SceneItemVisibilityChanged" event.
Since v4.0.0.
*/
type SceneItemVisibilityChanged struct {
	EventBasic

	// Scene item ID
	ItemId int `json:"item-id,omitempty"`

	// Name of the item in the scene.
	ItemName string `json:"item-name,omitempty"`

	// New visibility state of the item.
	ItemVisible bool `json:"item-visible"`

	// Name of the scene.
	SceneName string `json:"scene-name,omitempty"`
}
