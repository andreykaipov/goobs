// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
SceneListChanged represents the event body for the "SceneListChanged" event.
Since v5.0.0.
*/
type SceneListChanged struct {
	Scenes []*typedefs.Scene `json:"scenes,omitempty"`
}
