// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetStudioModeStatusParams represents the params body for the "GetStudioModeStatus" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStudioModeStatus.
*/
type GetStudioModeStatusParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetStudioModeStatusParams
func (o *GetStudioModeStatusParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetStudioModeStatusParams
func (o *GetStudioModeStatusParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetStudioModeStatusParams
func (o *GetStudioModeStatusParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetStudioModeStatusResponse represents the response body for the "GetStudioModeStatus" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStudioModeStatus.
*/
type GetStudioModeStatusResponse struct {
	requests.Response

	// Indicates if Studio Mode is enabled.
	StudioMode bool `json:"studio-mode"`
}

// GetMessageID returns the MessageID of GetStudioModeStatusResponse
func (o *GetStudioModeStatusResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetStudioModeStatusResponse
func (o *GetStudioModeStatusResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetStudioModeStatusResponse
func (o *GetStudioModeStatusResponse) GetError() string {
	return o.Error
}

// GetStudioModeStatus sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetStudioModeStatus(
	paramss ...*GetStudioModeStatusParams,
) (*GetStudioModeStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStudioModeStatusParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetStudioModeStatus"
	data := &GetStudioModeStatusResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
GetPreviewSceneParams represents the params body for the "GetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetPreviewScene.
*/
type GetPreviewSceneParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetPreviewSceneParams
func (o *GetPreviewSceneParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetPreviewSceneParams
func (o *GetPreviewSceneParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetPreviewSceneParams
func (o *GetPreviewSceneParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetPreviewSceneResponse represents the response body for the "GetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetPreviewScene.
*/
type GetPreviewSceneResponse struct {
	requests.Response

	// The name of the active preview scene.
	Name string `json:"name"`

	Sources []map[string]interface{} `json:"sources"`
}

// GetMessageID returns the MessageID of GetPreviewSceneResponse
func (o *GetPreviewSceneResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetPreviewSceneResponse
func (o *GetPreviewSceneResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetPreviewSceneResponse
func (o *GetPreviewSceneResponse) GetError() string {
	return o.Error
}

// GetPreviewScene sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetPreviewScene(
	paramss ...*GetPreviewSceneParams,
) (*GetPreviewSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetPreviewSceneParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetPreviewScene"
	data := &GetPreviewSceneResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
SetPreviewSceneParams represents the params body for the "SetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetPreviewScene.
*/
type SetPreviewSceneParams struct {
	requests.Params

	// The name of the scene to preview.
	SceneName string `json:"scene-name"`
}

// GetRequestType returns the RequestType of SetPreviewSceneParams
func (o *SetPreviewSceneParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetPreviewSceneParams
func (o *SetPreviewSceneParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetPreviewSceneParams
func (o *SetPreviewSceneParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetPreviewSceneResponse represents the response body for the "SetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetPreviewScene.
*/
type SetPreviewSceneResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetPreviewSceneResponse
func (o *SetPreviewSceneResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetPreviewSceneResponse
func (o *SetPreviewSceneResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetPreviewSceneResponse
func (o *SetPreviewSceneResponse) GetError() string {
	return o.Error
}

// SetPreviewScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetPreviewScene(params *SetPreviewSceneParams) (*SetPreviewSceneResponse, error) {
	params.RequestType = "SetPreviewScene"
	data := &SetPreviewSceneResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
TransitionToProgramParams represents the params body for the "TransitionToProgram" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#TransitionToProgram.
*/
type TransitionToProgramParams struct {
	requests.Params

	WithTransition struct {
		// Transition duration (in milliseconds).
		Duration int `json:"duration"`

		// Name of the transition.
		Name string `json:"name"`
	} `json:"with-transition"`
}

// GetRequestType returns the RequestType of TransitionToProgramParams
func (o *TransitionToProgramParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of TransitionToProgramParams
func (o *TransitionToProgramParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on TransitionToProgramParams
func (o *TransitionToProgramParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
TransitionToProgramResponse represents the response body for the "TransitionToProgram" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#TransitionToProgram.
*/
type TransitionToProgramResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of TransitionToProgramResponse
func (o *TransitionToProgramResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of TransitionToProgramResponse
func (o *TransitionToProgramResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of TransitionToProgramResponse
func (o *TransitionToProgramResponse) GetError() string {
	return o.Error
}

// TransitionToProgram sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) TransitionToProgram(
	params *TransitionToProgramParams,
) (*TransitionToProgramResponse, error) {
	params.RequestType = "TransitionToProgram"
	data := &TransitionToProgramResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
EnableStudioModeParams represents the params body for the "EnableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#EnableStudioMode.
*/
type EnableStudioModeParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of EnableStudioModeParams
func (o *EnableStudioModeParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of EnableStudioModeParams
func (o *EnableStudioModeParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on EnableStudioModeParams
func (o *EnableStudioModeParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
EnableStudioModeResponse represents the response body for the "EnableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#EnableStudioMode.
*/
type EnableStudioModeResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of EnableStudioModeResponse
func (o *EnableStudioModeResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of EnableStudioModeResponse
func (o *EnableStudioModeResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of EnableStudioModeResponse
func (o *EnableStudioModeResponse) GetError() string {
	return o.Error
}

// EnableStudioMode sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) EnableStudioMode(
	paramss ...*EnableStudioModeParams,
) (*EnableStudioModeResponse, error) {
	if len(paramss) == 0 {
		paramss = []*EnableStudioModeParams{{}}
	}
	params := paramss[0]
	params.RequestType = "EnableStudioMode"
	data := &EnableStudioModeResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
DisableStudioModeParams represents the params body for the "DisableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DisableStudioMode.
*/
type DisableStudioModeParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of DisableStudioModeParams
func (o *DisableStudioModeParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of DisableStudioModeParams
func (o *DisableStudioModeParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on DisableStudioModeParams
func (o *DisableStudioModeParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
DisableStudioModeResponse represents the response body for the "DisableStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DisableStudioMode.
*/
type DisableStudioModeResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of DisableStudioModeResponse
func (o *DisableStudioModeResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of DisableStudioModeResponse
func (o *DisableStudioModeResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of DisableStudioModeResponse
func (o *DisableStudioModeResponse) GetError() string {
	return o.Error
}

// DisableStudioMode sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) DisableStudioMode(
	paramss ...*DisableStudioModeParams,
) (*DisableStudioModeResponse, error) {
	if len(paramss) == 0 {
		paramss = []*DisableStudioModeParams{{}}
	}
	params := paramss[0]
	params.RequestType = "DisableStudioMode"
	data := &DisableStudioModeResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
ToggleStudioModeParams represents the params body for the "ToggleStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleStudioMode.
*/
type ToggleStudioModeParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of ToggleStudioModeParams
func (o *ToggleStudioModeParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of ToggleStudioModeParams
func (o *ToggleStudioModeParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on ToggleStudioModeParams
func (o *ToggleStudioModeParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ToggleStudioModeResponse represents the response body for the "ToggleStudioMode" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleStudioMode.
*/
type ToggleStudioModeResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of ToggleStudioModeResponse
func (o *ToggleStudioModeResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of ToggleStudioModeResponse
func (o *ToggleStudioModeResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of ToggleStudioModeResponse
func (o *ToggleStudioModeResponse) GetError() string {
	return o.Error
}

// ToggleStudioMode sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) ToggleStudioMode(
	paramss ...*ToggleStudioModeParams,
) (*ToggleStudioModeResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleStudioModeParams{{}}
	}
	params := paramss[0]
	params.RequestType = "ToggleStudioMode"
	data := &ToggleStudioModeResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
