// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the SceneItemCreated event.

A scene item has been created.
*/
type SceneItemCreated struct {
	// Numeric ID of the scene item
	SceneItemId int `json:"sceneItemId,omitempty"`

	// Index position of the item
	SceneItemIndex int `json:"sceneItemIndex,omitempty"`

	// Name of the scene the item was added to
	SceneName string `json:"sceneName,omitempty"`

	// UUID of the scene the item was added to
	SceneUuid string `json:"sceneUuid,omitempty"`

	// Name of the underlying source (input/scene)
	SourceName string `json:"sourceName,omitempty"`

	// UUID of the underlying source (input/scene)
	SourceUuid string `json:"sourceUuid,omitempty"`
}
