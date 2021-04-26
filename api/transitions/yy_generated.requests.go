// This file has been automatically generated. Don't edit it.

package transitions

/*
GetTransitionListParams represents the params body for the "GetTransitionList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionList.
*/
type GetTransitionListParams struct{}

/*
GetTransitionListResponse represents the response body for the "GetTransitionList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionList.
*/
type GetTransitionListResponse struct {
	// Name of the currently active transition.
	CurrentTransition string `json:"current-transition"`

	Transitions []struct {
		// Name of the transition.
		Name string `json:"name"`
	} `json:"transitions"`
}

/*
GetCurrentTransitionParams represents the params body for the "GetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentTransition.
*/
type GetCurrentTransitionParams struct{}

/*
GetCurrentTransitionResponse represents the response body for the "GetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentTransition.
*/
type GetCurrentTransitionResponse struct {
	// Transition duration (in milliseconds) if supported by the transition.
	Duration int `json:"duration"`

	// Name of the selected transition.
	Name string `json:"name"`
}

/*
SetCurrentTransitionParams represents the params body for the "SetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentTransition.
*/
type SetCurrentTransitionParams struct {
	// The name of the transition.
	TransitionName string `json:"transition-name"`
}

/*
SetCurrentTransitionResponse represents the response body for the "SetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentTransition.
*/
type SetCurrentTransitionResponse struct{}

/*
SetTransitionDurationParams represents the params body for the "SetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTransitionDuration.
*/
type SetTransitionDurationParams struct {
	// Desired duration of the transition (in milliseconds).
	Duration int `json:"duration"`
}

/*
SetTransitionDurationResponse represents the response body for the "SetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTransitionDuration.
*/
type SetTransitionDurationResponse struct{}

/*
GetTransitionDurationParams represents the params body for the "GetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionDuration.
*/
type GetTransitionDurationParams struct{}

/*
GetTransitionDurationResponse represents the response body for the "GetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionDuration.
*/
type GetTransitionDurationResponse struct {
	// Duration of the current transition (in milliseconds).
	TransitionDuration int `json:"transition-duration"`
}
