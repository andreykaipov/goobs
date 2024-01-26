// This file has been automatically generated. Don't edit it.

package events

/*
Represents the event body for the SceneItemRemoved event.

A scene item has been removed.

This event is not emitted when the scene the item is in is removed.
*/
type SceneItemRemoved struct {
	// Numeric ID of the scene item
	SceneItemId int `json:"sceneItemId,omitempty"`

	// Name of the scene the item was removed from
	SceneName string `json:"sceneName,omitempty"`

	// UUID of the scene the item was removed from
	SceneUuid string `json:"sceneUuid,omitempty"`

	// Name of the underlying source (input/scene)
	SourceName string `json:"sourceName,omitempty"`

	// UUID of the underlying source (input/scene)
	SourceUuid string `json:"sourceUuid,omitempty"`
}
