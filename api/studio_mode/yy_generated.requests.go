// This file has been automatically generated. Don't edit it.

package studiomode

type GetStudioModeStatusParams struct{}
type GetStudioModeStatusResponse struct {
	StudioMode bool `json:"studio-mode"`
}
type GetPreviewSceneParams struct{}
type GetPreviewSceneResponse struct {
	Name    string                   `json:"name"`
	Sources []map[string]interface{} `json:"sources"`
}
type SetPreviewSceneParams struct {
	SceneName string `json:"scene-name"`
}
type SetPreviewSceneResponse struct{}
type TransitionToProgramParams struct {
	WithTransition struct {
		Duration int    `json:"duration"`
		Name     string `json:"name"`
	}
}
type TransitionToProgramResponse struct{}
type EnableStudioModeParams struct{}
type EnableStudioModeResponse struct{}
type DisableStudioModeParams struct{}
type DisableStudioModeResponse struct{}
type ToggleStudioModeParams struct{}
type ToggleStudioModeResponse struct{}
