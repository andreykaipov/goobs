// This file has been automatically generated. Don't edit it.

package studiomode

// GetStudioModeStatusParams contains the request body for the
// [GetStudioModeStatus](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStudioModeStatus)
// request.
type GetStudioModeStatusParams struct{}

// GetStudioModeStatusResponse contains the request body for the
// [GetStudioModeStatus](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStudioModeStatus)
// request.
type GetStudioModeStatusResponse struct {
	// Indicates if Studio Mode is enabled.
	StudioMode bool `json:"studio-mode"`
}

// GetPreviewSceneParams contains the request body for the
// [GetPreviewScene](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetPreviewScene)
// request.
type GetPreviewSceneParams struct{}

// GetPreviewSceneResponse contains the request body for the
// [GetPreviewScene](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetPreviewScene)
// request.
type GetPreviewSceneResponse struct {
	// The name of the active preview scene.
	Name    string                   `json:"name"`
	Sources []map[string]interface{} `json:"sources"`
}

// SetPreviewSceneParams contains the request body for the
// [SetPreviewScene](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetPreviewScene)
// request.
type SetPreviewSceneParams struct {
	// The name of the scene to preview.
	SceneName string `json:"scene-name"`
}

// SetPreviewSceneResponse contains the request body for the
// [SetPreviewScene](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetPreviewScene)
// request.
type SetPreviewSceneResponse struct{}

// TransitionToProgramParams contains the request body for the
// [TransitionToProgram](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#TransitionToProgram)
// request.
type TransitionToProgramParams struct {
	WithTransition struct {
		// Transition duration (in milliseconds).
		Duration int `json:"duration"`
		// Name of the transition.
		Name string `json:"name"`
	} `json:"with-transition"`
}

// TransitionToProgramResponse contains the request body for the
// [TransitionToProgram](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#TransitionToProgram)
// request.
type TransitionToProgramResponse struct{}

// EnableStudioModeParams contains the request body for the
// [EnableStudioMode](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#EnableStudioMode)
// request.
type EnableStudioModeParams struct{}

// EnableStudioModeResponse contains the request body for the
// [EnableStudioMode](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#EnableStudioMode)
// request.
type EnableStudioModeResponse struct{}

// DisableStudioModeParams contains the request body for the
// [DisableStudioMode](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DisableStudioMode)
// request.
type DisableStudioModeParams struct{}

// DisableStudioModeResponse contains the request body for the
// [DisableStudioMode](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DisableStudioMode)
// request.
type DisableStudioModeResponse struct{}

// ToggleStudioModeParams contains the request body for the
// [ToggleStudioMode](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleStudioMode)
// request.
type ToggleStudioModeParams struct{}

// ToggleStudioModeResponse contains the request body for the
// [ToggleStudioMode](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleStudioMode)
// request.
type ToggleStudioModeResponse struct{}
