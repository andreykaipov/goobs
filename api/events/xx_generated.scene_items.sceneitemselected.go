// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the SceneItemSelected event.

A scene item has been selected in the Ui.
*/
type SceneItemSelected struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}
