// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
PreviewSceneChanged represents the event body for the "PreviewSceneChanged" event.
Since v4.1.0.
*/
type PreviewSceneChanged struct {
	EventBasic

	// Name of the scene being previewed.
	SceneName string `json:"scene-name,omitempty"`

	// List of sources composing the scene. Same specification as [`GetCurrentScene`](#getcurrentscene).
	Sources []typedefs.SceneItem `json:"sources,omitempty"`
}
