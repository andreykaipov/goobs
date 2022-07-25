// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemListReindexed represents the event body for the "SceneItemListReindexed" event.
Since v5.0.0.
*/
type SceneItemListReindexed struct {
	// Array of scene item objects
	SceneItems []interface{} `json:"sceneItems,omitempty"`

	// Name of the scene
	SceneName string `json:"sceneName,omitempty"`
}
