// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemSelected represents the event body for the "SceneItemSelected" event.
Since v5.0.0.
*/
type SceneItemSelected struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}
