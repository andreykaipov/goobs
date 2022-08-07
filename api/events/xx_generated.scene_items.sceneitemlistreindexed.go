// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
SceneItemListReindexed represents the event body for the "SceneItemListReindexed" event.
Since v5.0.0.
*/
type SceneItemListReindexed struct {
	SceneItems []*typedefs.SceneItem `json:"sceneItems,omitempty"`

	// Name of the scene
	SceneName string `json:"sceneName,omitempty"`
}
