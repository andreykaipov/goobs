// This file has been automatically generated. Don't edit it.

package transitions

// GetTransitionListParams contains the request body for the
// [GetTransitionList](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionList)
// request.
type GetTransitionListParams struct{}

// GetTransitionListResponse contains the request body for the
// [GetTransitionList](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionList)
// request.
type GetTransitionListResponse struct {
	// Name of the currently active transition.
	CurrentTransition string `json:"current-transition"`
	Transitions       []struct {
		// Name of the transition.
		Name string `json:"name"`
	} `json:"transitions"`
}

// GetCurrentTransitionParams contains the request body for the
// [GetCurrentTransition](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentTransition)
// request.
type GetCurrentTransitionParams struct{}

// GetCurrentTransitionResponse contains the request body for the
// [GetCurrentTransition](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentTransition)
// request.
type GetCurrentTransitionResponse struct {
	// Transition duration (in milliseconds) if supported by the transition.
	Duration int `json:"duration"`
	// Name of the selected transition.
	Name string `json:"name"`
}

// SetCurrentTransitionParams contains the request body for the
// [SetCurrentTransition](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentTransition)
// request.
type SetCurrentTransitionParams struct {
	// The name of the transition.
	TransitionName string `json:"transition-name"`
}

// SetCurrentTransitionResponse contains the request body for the
// [SetCurrentTransition](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentTransition)
// request.
type SetCurrentTransitionResponse struct{}

// SetTransitionDurationParams contains the request body for the
// [SetTransitionDuration](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTransitionDuration)
// request.
type SetTransitionDurationParams struct {
	// Desired duration of the transition (in milliseconds).
	Duration int `json:"duration"`
}

// SetTransitionDurationResponse contains the request body for the
// [SetTransitionDuration](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTransitionDuration)
// request.
type SetTransitionDurationResponse struct{}

// GetTransitionDurationResponse contains the request body for the
// [GetTransitionDuration](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionDuration)
// request.
type GetTransitionDurationResponse struct {
	// Duration of the current transition (in milliseconds).
	TransitionDuration int `json:"transition-duration"`
}

// GetTransitionDurationParams contains the request body for the
// [GetTransitionDuration](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionDuration)
// request.
type GetTransitionDurationParams struct{}
