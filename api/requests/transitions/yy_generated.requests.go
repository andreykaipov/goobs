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

// GetRequestType returns the RequestType of GetTransitionListParams
func (o *GetTransitionListParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetTransitionListParams
func (o *GetTransitionListParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetTransitionListParams
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

// GetMessageID returns the MessageID of GetTransitionListResponse
func (o *GetTransitionListResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetTransitionListResponse
func (o *GetTransitionListResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetTransitionListResponse
func (o *GetTransitionListResponse) GetError() string {
	return o.Error
}

// GetTransitionList sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetTransitionList(
	paramss ...*GetTransitionListParams,
) (*GetTransitionListResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetTransitionListParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetTransitionList"
	data := &GetTransitionListResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
GetCurrentTransitionParams represents the params body for the "GetCurrentTransition" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetCurrentTransition.
*/
type GetCurrentTransitionParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetCurrentTransitionParams
func (o *GetCurrentTransitionParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetCurrentTransitionParams
func (o *GetCurrentTransitionParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetCurrentTransitionParams
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

// GetMessageID returns the MessageID of GetCurrentTransitionResponse
func (o *GetCurrentTransitionResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetCurrentTransitionResponse
func (o *GetCurrentTransitionResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetCurrentTransitionResponse
func (o *GetCurrentTransitionResponse) GetError() string {
	return o.Error
}

// GetCurrentTransition sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetCurrentTransition(
	paramss ...*GetCurrentTransitionParams,
) (*GetCurrentTransitionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetCurrentTransitionParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetCurrentTransition"
	data := &GetCurrentTransitionResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
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

// GetRequestType returns the RequestType of SetCurrentTransitionParams
func (o *SetCurrentTransitionParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetCurrentTransitionParams
func (o *SetCurrentTransitionParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetCurrentTransitionParams
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

// GetMessageID returns the MessageID of SetCurrentTransitionResponse
func (o *SetCurrentTransitionResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetCurrentTransitionResponse
func (o *SetCurrentTransitionResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetCurrentTransitionResponse
func (o *SetCurrentTransitionResponse) GetError() string {
	return o.Error
}

// SetCurrentTransition sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentTransition(
	params *SetCurrentTransitionParams,
) (*SetCurrentTransitionResponse, error) {
	params.RequestType = "SetCurrentTransition"
	data := &SetCurrentTransitionResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
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

// GetRequestType returns the RequestType of SetTransitionDurationParams
func (o *SetTransitionDurationParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetTransitionDurationParams
func (o *SetTransitionDurationParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetTransitionDurationParams
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

// GetMessageID returns the MessageID of SetTransitionDurationResponse
func (o *SetTransitionDurationResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetTransitionDurationResponse
func (o *SetTransitionDurationResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetTransitionDurationResponse
func (o *SetTransitionDurationResponse) GetError() string {
	return o.Error
}

// SetTransitionDuration sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetTransitionDuration(
	params *SetTransitionDurationParams,
) (*SetTransitionDurationResponse, error) {
	params.RequestType = "SetTransitionDuration"
	data := &SetTransitionDurationResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
GetTransitionDurationParams represents the params body for the "GetTransitionDuration" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetTransitionDuration.
*/
type GetTransitionDurationParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetTransitionDurationParams
func (o *GetTransitionDurationParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetTransitionDurationParams
func (o *GetTransitionDurationParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetTransitionDurationParams
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

// GetMessageID returns the MessageID of GetTransitionDurationResponse
func (o *GetTransitionDurationResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetTransitionDurationResponse
func (o *GetTransitionDurationResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetTransitionDurationResponse
func (o *GetTransitionDurationResponse) GetError() string {
	return o.Error
}

// GetTransitionDuration sends the corresponding request to the connected OBS WebSockets server.
// Note the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetTransitionDuration(
	paramss ...*GetTransitionDurationParams,
) (*GetTransitionDurationResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetTransitionDurationParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetTransitionDuration"
	data := &GetTransitionDurationResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
