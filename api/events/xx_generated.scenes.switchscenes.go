// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
SwitchScenes represents the event body for the "SwitchScenes" event.
Since v0.3.
*/
type SwitchScenes struct {
	EventBasic

	// The new scene.
	SceneName string `json:"scene-name,omitempty"`

	// List of scene items in the new scene. Same specification as [`GetCurrentScene`](#getcurrentscene).
	Sources []typedefs.SceneItem `json:"sources,omitempty"`
}
