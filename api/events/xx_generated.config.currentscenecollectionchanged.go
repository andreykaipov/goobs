// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the CurrentSceneCollectionChanged event.

The current scene collection has changed.

Note: If polling has been paused during `CurrentSceneCollectionChanging`, this is the que to restart polling.
*/
type CurrentSceneCollectionChanged struct {
	// Name of the new scene collection
	SceneCollectionName string `json:"sceneCollectionName,omitempty"`
}
