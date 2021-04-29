// This file has been automatically generated. Don't edit it.

package events

/*
PreviewSceneChanged represents the event body for the "PreviewSceneChanged" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#PreviewSceneChanged.
*/
type PreviewSceneChanged struct {
	EventBasic

	// Name of the scene being previewed.
	SceneName string `json:"scene-name"`

	// List of sources composing the scene. Same specification as
	// [`GetCurrentScene`](#getcurrentscene).
	Sources []map[string]interface{} `json:"sources"`
}
