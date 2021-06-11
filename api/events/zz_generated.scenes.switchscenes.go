// This file has been automatically generated. Don't edit it.

package events

/*
SwitchScenes represents the event body for the "SwitchScenes" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#SwitchScenes.
*/
type SwitchScenes struct {
	EventBasic

	// The new scene.
	SceneName string `json:"scene-name"`

	// List of sources in the new scene. Same specification as
	// [`GetCurrentScene`](#getcurrentscene).
	Sources []map[string]interface{} `json:"sources"`
}
