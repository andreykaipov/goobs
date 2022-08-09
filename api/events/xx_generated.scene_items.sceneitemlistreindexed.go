// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
Represents the event body for the SceneItemListReindexed event.

A scene's item list has been reindexed.
*/
type SceneItemListReindexed struct {
	SceneItems []*typedefs.SceneItem `json:"sceneItems,omitempty"`

	// Name of the scene
	SceneName string `json:"sceneName,omitempty"`
}
