// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
SceneItemTransformChanged represents the event body for the "SceneItemTransformChanged" event.
Since v4.6.0.
*/
type SceneItemTransformChanged struct {
	EventBasic

	// Scene item ID
	ItemId int `json:"item-id,omitempty"`

	// Name of the item in the scene.
	ItemName string `json:"item-name,omitempty"`

	// Name of the scene.
	SceneName string `json:"scene-name,omitempty"`

	// Scene item transform properties
	Transform typedefs.SceneItemTransform `json:"transform,omitempty"`
}
