// This file has been automatically generated. Don't edit it.

package transitions

type GetTransitionListParams struct{}
type GetTransitionListResponse struct {
	// Name of the currently active transition.
	CurrentTransition string `json:"current-transition"`
	Transitions       []struct {
		// Name of the transition.
		Name string `json:"name"`
	} `json:"transitions"`
}
type GetCurrentTransitionParams struct{}
type GetCurrentTransitionResponse struct {
	// Transition duration (in milliseconds) if supported by the transition.
	Duration int `json:"duration"`
	// Name of the selected transition.
	Name string `json:"name"`
}
type SetCurrentTransitionParams struct {
	// The name of the transition.
	TransitionName string `json:"transition-name"`
}
type SetCurrentTransitionResponse struct{}
type SetTransitionDurationParams struct {
	// Desired duration of the transition (in milliseconds).
	Duration int `json:"duration"`
}
type SetTransitionDurationResponse struct{}
type GetTransitionDurationParams struct{}
type GetTransitionDurationResponse struct {
	// Duration of the current transition (in milliseconds).
	TransitionDuration int `json:"transition-duration"`
}
