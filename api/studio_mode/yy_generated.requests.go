// This file has been automatically generated. Don't edit it.

package studiomode

type GetStudioModeStatusParams struct{}
type GetStudioModeStatusResponse struct {
	// Indicates if Studio Mode is enabled.
	StudioMode bool `json:"studio-mode"`
}
type GetPreviewSceneParams struct{}
type GetPreviewSceneResponse struct {
	// The name of the active preview scene.
	Name    string                   `json:"name"`
	Sources []map[string]interface{} `json:"sources"`
}
type SetPreviewSceneParams struct {
	// The name of the scene to preview.
	SceneName string `json:"scene-name"`
}
type SetPreviewSceneResponse struct{}
type TransitionToProgramParams struct {
	WithTransition struct {
		// Transition duration (in milliseconds).
		Duration int `json:"duration"`
		// Name of the transition.
		Name string `json:"name"`
	} `json:"with-transition"`
}
type TransitionToProgramResponse struct{}
type EnableStudioModeParams struct{}
type EnableStudioModeResponse struct{}
type DisableStudioModeParams struct{}
type DisableStudioModeResponse struct{}
type ToggleStudioModeParams struct{}
type ToggleStudioModeResponse struct{}
