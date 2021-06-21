// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemRemoved represents the event body for the "SceneItemRemoved" event.
Since v4.0.0.
*/
type SceneItemRemoved struct {
	EventBasic

	// Scene item ID
	ItemId int `json:"item-id,omitempty"`

	// Name of the item removed from the scene.
	ItemName string `json:"item-name,omitempty"`

	// Name of the scene.
	SceneName string `json:"scene-name,omitempty"`
}
