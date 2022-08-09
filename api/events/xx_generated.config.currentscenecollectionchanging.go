// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the CurrentSceneCollectionChanging event.

The current scene collection has begun changing.

Note: We recommend using this event to trigger a pause of all polling requests, as performing any requests during a
scene collection change is considered undefined behavior and can cause crashes!
*/
type CurrentSceneCollectionChanging struct {
	// Name of the current scene collection
	SceneCollectionName string `json:"sceneCollectionName,omitempty"`
}
