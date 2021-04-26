// This file has been automatically generated. Don't edit it.

package studiomode

import api "github.com/andreykaipov/goobs/api"

/*
GetStudioModeStatusParams represents the params body for the "GetStudioModeStatus" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStudioModeStatus.
*/
type GetStudioModeStatusParams struct {
	api.Params
}

func (o *GetStudioModeStatusParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetStudioModeStatusParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetStudioModeStatusResponse represents the response body for the "GetStudioModeStatus" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStudioModeStatus.
*/
type GetStudioModeStatusResponse struct {
	api.Response

	// Indicates if Studio Mode is enabled.
	StudioMode bool `json:"studio-mode"`
}

func (o *GetStudioModeStatusResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetStudioModeStatusResponse) GetStatus() string {
	return o.Status
}
func (o *GetStudioModeStatusResponse) GetError() string {
	return o.Error
}

/*
GetPreviewSceneParams represents the params body for the "GetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetPreviewScene.
*/
type GetPreviewSceneParams struct {
	api.Params
}

func (o *GetPreviewSceneParams) GetRequestType() string {
	return o.RequestType
}
func (o *GetPreviewSceneParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetPreviewSceneResponse represents the response body for the "GetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetPreviewScene.
*/
type GetPreviewSceneResponse struct {
	api.Response

	// The name of the active preview scene.
	Name string `json:"name"`

	Sources []map[string]interface{} `json:"sources"`
}

func (o *GetPreviewSceneResponse) GetMessageID() string {
	return o.MessageID
}
func (o *GetPreviewSceneResponse) GetStatus() string {
	return o.Status
}
func (o *GetPreviewSceneResponse) GetError() string {
	return o.Error
}

/*
SetPreviewSceneParams represents the params body for the "SetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetPreviewScene.
*/
type SetPreviewSceneParams struct {
	api.Params

	// The name of the scene to preview.
	SceneName string `json:"scene-name"`
}

func (o *SetPreviewSceneParams) GetRequestType() string {
	return o.RequestType
}
func (o *SetPreviewSceneParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetPreviewSceneResponse represents the response body for the "SetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetPreviewScene.
*/
type SetPreviewSceneResponse struct {
	api.Response
}

func (o *SetPreviewSceneResponse) GetMessageID() string {
	return o.MessageID
}
func (o *SetPreviewSceneResponse) GetStatus() string {
	return o.Status
}
func (o *SetPreviewSceneResponse) GetError() string {
	return o.Error
}

/*
TransitionToProgramParams represents the params body for the "TransitionToProgram" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#TransitionToProgram.
*/
type TransitionToProgramParams struct {
	api.Params

	WithTransition struct {
		// Transition duration (in milliseconds).
		Duration int `json:"duration"`

		// Name of the transition.
		Name string `json:"name"`
	} `json:"with-transition"`
}

func (o *TransitionToProgramParams) GetRequestType() string {
	return o.RequestType
}
func (o *TransitionToProgramParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
TransitionToProgramResponse represents the response body for the "TransitionToProgram" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#TransitionToProgram.
*/
type TransitionToProgramResponse struct {
	api.Response
}

func (o *TransitionToProgramResponse) GetMessageID() string {
	return o.MessageID
}
func (o *TransitionToProgramResponse) GetStatus() string {
	return o.Status
}
func (o *TransitionToProgramResponse) GetError() string {
	return o.Error
}

/*
EnableStudioModeParams represents the params body for the "EnableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#EnableStudioMode.
*/
type EnableStudioModeParams struct {
	api.Params
}

func (o *EnableStudioModeParams) GetRequestType() string {
	return o.RequestType
}
func (o *EnableStudioModeParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
EnableStudioModeResponse represents the response body for the "EnableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#EnableStudioMode.
*/
type EnableStudioModeResponse struct {
	api.Response
}

func (o *EnableStudioModeResponse) GetMessageID() string {
	return o.MessageID
}
func (o *EnableStudioModeResponse) GetStatus() string {
	return o.Status
}
func (o *EnableStudioModeResponse) GetError() string {
	return o.Error
}

/*
DisableStudioModeParams represents the params body for the "DisableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DisableStudioMode.
*/
type DisableStudioModeParams struct {
	api.Params
}

func (o *DisableStudioModeParams) GetRequestType() string {
	return o.RequestType
}
func (o *DisableStudioModeParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
DisableStudioModeResponse represents the response body for the "DisableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DisableStudioMode.
*/
type DisableStudioModeResponse struct {
	api.Response
}

func (o *DisableStudioModeResponse) GetMessageID() string {
	return o.MessageID
}
func (o *DisableStudioModeResponse) GetStatus() string {
	return o.Status
}
func (o *DisableStudioModeResponse) GetError() string {
	return o.Error
}

/*
ToggleStudioModeParams represents the params body for the "ToggleStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleStudioMode.
*/
type ToggleStudioModeParams struct {
	api.Params
}

func (o *ToggleStudioModeParams) GetRequestType() string {
	return o.RequestType
}
func (o *ToggleStudioModeParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ToggleStudioModeResponse represents the response body for the "ToggleStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleStudioMode.
*/
type ToggleStudioModeResponse struct {
	api.Response
}

func (o *ToggleStudioModeResponse) GetMessageID() string {
	return o.MessageID
}
func (o *ToggleStudioModeResponse) GetStatus() string {
	return o.Status
}
func (o *ToggleStudioModeResponse) GetError() string {
	return o.Error
}
