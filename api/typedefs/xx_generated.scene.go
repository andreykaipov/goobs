// This file has been automatically generated. Don't edit it.

package typedefs

// Scene represents the complex type for Scene.
type Scene struct {
	// Name of the currently active scene.
	Name string `json:"name,omitempty"`

	// Ordered list of the current scene's source items.
	Sources []SceneItem `json:"sources,omitempty"`
}
