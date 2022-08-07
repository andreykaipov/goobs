// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
SceneItemTransformChanged represents the event body for the "SceneItemTransformChanged" event.
Since v5.0.0.
*/
type SceneItemTransformChanged struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Scene item transform info
	SceneItemTransform *typedefs.SceneItemTransform `json:"sceneItemTransform,omitempty"`

	// The name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}
