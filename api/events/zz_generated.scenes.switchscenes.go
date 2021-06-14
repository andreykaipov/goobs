// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
SwitchScenes represents the event body for the "SwitchScenes" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#SwitchScenes.
*/
type SwitchScenes struct {
	EventBasic

	// The new scene.
	SceneName string `json:"scene-name"`

	// List of scene items in the new scene. Same specification as
	// [`GetCurrentScene`](#getcurrentscene).
	Sources []typedefs.SceneItem `json:"sources"`
}
