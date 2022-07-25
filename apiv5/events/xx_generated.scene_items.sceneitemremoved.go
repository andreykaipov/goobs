// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemRemoved represents the event body for the "SceneItemRemoved" event.
Since v5.0.0.
*/
type SceneItemRemoved struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item was removed from
	SceneName string `json:"sceneName,omitempty"`

	// Name of the underlying source (input/scene)
	SourceName string `json:"sourceName,omitempty"`
}
