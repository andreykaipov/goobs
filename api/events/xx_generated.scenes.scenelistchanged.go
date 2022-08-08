// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the event body for the SceneListChanged event.
type SceneListChanged struct {
	Scenes []*typedefs.Scene `json:"scenes,omitempty"`
}
