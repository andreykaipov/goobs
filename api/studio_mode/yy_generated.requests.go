// This file has been automatically generated. Don't edit it.

package studiomode

/*
GetStudioModeStatusParams represents the params body for the "GetStudioModeStatus" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStudioModeStatus.
*/
type GetStudioModeStatusParams struct{}

/*
GetStudioModeStatusResponse represents the response body for the "GetStudioModeStatus" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStudioModeStatus.
*/
type GetStudioModeStatusResponse struct {
	// Indicates if Studio Mode is enabled.
	StudioMode bool `json:"studio-mode"`
}

/*
GetPreviewSceneParams represents the params body for the "GetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetPreviewScene.
*/
type GetPreviewSceneParams struct{}

/*
GetPreviewSceneResponse represents the response body for the "GetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetPreviewScene.
*/
type GetPreviewSceneResponse struct {
	// The name of the active preview scene.
	Name string `json:"name"`

	Sources []map[string]interface{} `json:"sources"`
}

/*
SetPreviewSceneParams represents the params body for the "SetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetPreviewScene.
*/
type SetPreviewSceneParams struct {
	// The name of the scene to preview.
	SceneName string `json:"scene-name"`
}

/*
SetPreviewSceneResponse represents the response body for the "SetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetPreviewScene.
*/
type SetPreviewSceneResponse struct{}

/*
TransitionToProgramParams represents the params body for the "TransitionToProgram" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#TransitionToProgram.
*/
type TransitionToProgramParams struct {
	WithTransition struct {
		// Transition duration (in milliseconds).
		Duration int `json:"duration"`

		// Name of the transition.
		Name string `json:"name"`
	} `json:"with-transition"`
}

/*
TransitionToProgramResponse represents the response body for the "TransitionToProgram" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#TransitionToProgram.
*/
type TransitionToProgramResponse struct{}

/*
EnableStudioModeParams represents the params body for the "EnableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#EnableStudioMode.
*/
type EnableStudioModeParams struct{}

/*
EnableStudioModeResponse represents the response body for the "EnableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#EnableStudioMode.
*/
type EnableStudioModeResponse struct{}

/*
DisableStudioModeParams represents the params body for the "DisableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DisableStudioMode.
*/
type DisableStudioModeParams struct{}

/*
DisableStudioModeResponse represents the response body for the "DisableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DisableStudioMode.
*/
type DisableStudioModeResponse struct{}

/*
ToggleStudioModeParams represents the params body for the "ToggleStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleStudioMode.
*/
type ToggleStudioModeParams struct{}

/*
ToggleStudioModeResponse represents the response body for the "ToggleStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleStudioMode.
*/
type ToggleStudioModeResponse struct{}
