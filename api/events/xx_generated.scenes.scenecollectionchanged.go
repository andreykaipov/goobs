// This file has been automatically generated. Don't edit it.

package events

/*
SceneCollectionChanged represents the event body for the "SceneCollectionChanged" event.
Since v4.0.0.
*/
type SceneCollectionChanged struct {
	EventBasic

	// Name of the new current scene collection.
	SceneCollection string `json:"sceneCollection,omitempty"`
}
