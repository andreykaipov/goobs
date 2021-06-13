// This file has been automatically generated. Don't edit it.

package typedefs

/*
Scene represents the complex type for Scene.

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#Scene.
*/
type Scene struct {
	// Name of the currently active scene.
	Name string `json:"name"`

	// Ordered list of the current scene's source items.
	Sources []SceneItem `json:"sources"`
}
