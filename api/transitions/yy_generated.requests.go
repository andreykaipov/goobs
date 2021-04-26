// This file has been automatically generated. Don't edit it.

package transitions

type GetTransitionListParams struct{}
type GetTransitionListResponse struct {
	CurrentTransition string `json:"current-transition"`
	Transitions       []struct {
		Name string `json:"name"`
	}
}
type GetCurrentTransitionParams struct{}
type GetCurrentTransitionResponse struct {
	Duration int    `json:"duration"`
	Name     string `json:"name"`
}
type SetCurrentTransitionParams struct {
	TransitionName string `json:"transition-name"`
}
type SetCurrentTransitionResponse struct{}
type SetTransitionDurationParams struct {
	Duration int `json:"duration"`
}
type SetTransitionDurationResponse struct{}
type GetTransitionDurationParams struct{}
type GetTransitionDurationResponse struct {
	TransitionDuration int `json:"transition-duration"`
}
