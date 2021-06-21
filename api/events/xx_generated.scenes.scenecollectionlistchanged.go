// This file has been automatically generated. Don't edit it.

package events

/*
SceneCollectionListChanged represents the event body for the "SceneCollectionListChanged" event.
Since v4.0.0.
*/
type SceneCollectionListChanged struct {
	EventBasic

	SceneCollections []*SceneCollection `json:"sceneCollections,omitempty"`
}

type SceneCollection struct {
	// Scene collection name.
	Name string `json:"name,omitempty"`
}
