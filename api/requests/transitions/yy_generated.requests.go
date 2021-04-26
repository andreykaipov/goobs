// This file has been automatically generated. Don't edit it.

package transitions

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetTransitionListParams represents the params body for the "GetTransitionList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionList.
*/
type GetTransitionListParams struct {
	requests.Params
}

func (o *GetTransitionListParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetTransitionListParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetTransitionListResponse represents the response body for the "GetTransitionList" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionList.
*/
type GetTransitionListResponse struct {
	requests.Response

	// Name of the currently active transition.
	CurrentTransition string `json:"current-transition"`

	Transitions []struct {
		// Name of the transition.
		Name string `json:"name"`
	} `json:"transitions"`
}

func (o *GetTransitionListResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetTransitionListResponse) GetStatus() string {
	return o.Status
}
func (o *GetTransitionListResponse) GetError() string {
	return o.Error
}

/*
GetCurrentTransitionParams represents the params body for the "GetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentTransition.
*/
type GetCurrentTransitionParams struct {
	requests.Params
}

func (o *GetCurrentTransitionParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetCurrentTransitionParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetCurrentTransitionResponse represents the response body for the "GetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentTransition.
*/
type GetCurrentTransitionResponse struct {
	requests.Response

	// Transition duration (in milliseconds) if supported by the transition.
	Duration int `json:"duration"`

	// Name of the selected transition.
	Name string `json:"name"`
}

func (o *GetCurrentTransitionResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetCurrentTransitionResponse) GetStatus() string {
	return o.Status
}
func (o *GetCurrentTransitionResponse) GetError() string {
	return o.Error
}

/*
SetCurrentTransitionParams represents the params body for the "SetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentTransition.
*/
type SetCurrentTransitionParams struct {
	requests.Params

	// The name of the transition.
	TransitionName string `json:"transition-name"`
}

func (o *SetCurrentTransitionParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetCurrentTransitionParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetCurrentTransitionResponse represents the response body for the "SetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetCurrentTransition.
*/
type SetCurrentTransitionResponse struct {
	requests.Response
}

func (o *SetCurrentTransitionResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetCurrentTransitionResponse) GetStatus() string {
	return o.Status
}
func (o *SetCurrentTransitionResponse) GetError() string {
	return o.Error
}

/*
SetTransitionDurationParams represents the params body for the "SetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTransitionDuration.
*/
type SetTransitionDurationParams struct {
	requests.Params

	// Desired duration of the transition (in milliseconds).
	Duration int `json:"duration"`
}

func (o *SetTransitionDurationParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetTransitionDurationParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetTransitionDurationResponse represents the response body for the "SetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTransitionDuration.
*/
type SetTransitionDurationResponse struct {
	requests.Response
}

func (o *SetTransitionDurationResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetTransitionDurationResponse) GetStatus() string {
	return o.Status
}
func (o *SetTransitionDurationResponse) GetError() string {
	return o.Error
}

/*
GetTransitionDurationParams represents the params body for the "GetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionDuration.
*/
type GetTransitionDurationParams struct {
	requests.Params
}

func (o *GetTransitionDurationParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetTransitionDurationParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetTransitionDurationResponse represents the response body for the "GetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionDuration.
*/
type GetTransitionDurationResponse struct {
	requests.Response

	// Duration of the current transition (in milliseconds).
	TransitionDuration int `json:"transition-duration"`
}

func (o *GetTransitionDurationResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetTransitionDurationResponse) GetStatus() string {
	return o.Status
}
func (o *GetTransitionDurationResponse) GetError() string {
	return o.Error
}
