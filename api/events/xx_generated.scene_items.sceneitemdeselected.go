// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemDeselected represents the event body for the "SceneItemDeselected" event.
Since v4.6.0.
*/
type SceneItemDeselected struct {
	EventBasic

	// Name of the item in the scene.
	ItemId int `json:"item-id,omitempty"`

	// Name of the item in the scene.
	ItemName string `json:"item-name,omitempty"`

	// Name of the scene.
	SceneName string `json:"scene-name,omitempty"`
}
