// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetVersionParams represents the params body for the "GetVersion" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVersion.
*/
type GetVersionParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetVersionParams
func (o *GetVersionParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetVersionParams
func (o *GetVersionParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetVersionParams
func (o *GetVersionParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetVersionResponse represents the response body for the "GetVersion" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetVersion.
*/
type GetVersionResponse struct {
	requests.Response

	// List of available request types, formatted as a comma-separated list string (e.g. :
	// "Method1,Method2,Method3").
	AvailableRequests string `json:"available-requests"`

	// OBS Studio program version.
	ObsStudioVersion string `json:"obs-studio-version"`

	// obs-websocket plugin version.
	ObsWebsocketVersion string `json:"obs-websocket-version"`

	// OBSRemote compatible API version. Fixed to 1.1 for retrocompatibility.
	Version float64 `json:"version"`
}

// GetMessageID returns the MessageID of GetVersionResponse
func (o *GetVersionResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetVersionResponse
func (o *GetVersionResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetVersionResponse
func (o *GetVersionResponse) GetError() string {
	return o.Error
}

// GetVersion sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetVersion(paramss ...*GetVersionParams) (*GetVersionResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetVersionParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetVersion"
	data := &GetVersionResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
GetAuthRequiredParams represents the params body for the "GetAuthRequired" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetAuthRequired.
*/
type GetAuthRequiredParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetAuthRequiredParams
func (o *GetAuthRequiredParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetAuthRequiredParams
func (o *GetAuthRequiredParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetAuthRequiredParams
func (o *GetAuthRequiredParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetAuthRequiredResponse represents the response body for the "GetAuthRequired" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetAuthRequired.
*/
type GetAuthRequiredResponse struct {
	requests.Response

	// Indicates whether authentication is required.
	AuthRequired bool `json:"authRequired"`

	Challenge string `json:"challenge"`

	Salt string `json:"salt"`
}

// GetMessageID returns the MessageID of GetAuthRequiredResponse
func (o *GetAuthRequiredResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetAuthRequiredResponse
func (o *GetAuthRequiredResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetAuthRequiredResponse
func (o *GetAuthRequiredResponse) GetError() string {
	return o.Error
}

// GetAuthRequired sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) GetAuthRequired(
	paramss ...*GetAuthRequiredParams,
) (*GetAuthRequiredResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetAuthRequiredParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetAuthRequired"
	data := &GetAuthRequiredResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
AuthenticateParams represents the params body for the "Authenticate" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#Authenticate.
*/
type AuthenticateParams struct {
	requests.Params

	// Response to the auth challenge (see "Authentication" for more information).
	Auth string `json:"auth"`
}

// GetRequestType returns the RequestType of AuthenticateParams
func (o *AuthenticateParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of AuthenticateParams
func (o *AuthenticateParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on AuthenticateParams
func (o *AuthenticateParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
AuthenticateResponse represents the response body for the "Authenticate" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#Authenticate.
*/
type AuthenticateResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of AuthenticateResponse
func (o *AuthenticateResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of AuthenticateResponse
func (o *AuthenticateResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of AuthenticateResponse
func (o *AuthenticateResponse) GetError() string {
	return o.Error
}

// Authenticate sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) Authenticate(params *AuthenticateParams) (*AuthenticateResponse, error) {
	params.RequestType = "Authenticate"
	data := &AuthenticateResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
SetHeartbeatParams represents the params body for the "SetHeartbeat" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetHeartbeat.
*/
type SetHeartbeatParams struct {
	requests.Params

	// Starts/Stops emitting heartbeat messages
	Enable bool `json:"enable"`
}

// GetRequestType returns the RequestType of SetHeartbeatParams
func (o *SetHeartbeatParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetHeartbeatParams
func (o *SetHeartbeatParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetHeartbeatParams
func (o *SetHeartbeatParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetHeartbeatResponse represents the response body for the "SetHeartbeat" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetHeartbeat.
*/
type SetHeartbeatResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetHeartbeatResponse
func (o *SetHeartbeatResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetHeartbeatResponse
func (o *SetHeartbeatResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetHeartbeatResponse
func (o *SetHeartbeatResponse) GetError() string {
	return o.Error
}

// SetHeartbeat sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetHeartbeat(params *SetHeartbeatParams) (*SetHeartbeatResponse, error) {
	params.RequestType = "SetHeartbeat"
	data := &SetHeartbeatResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
SetFilenameFormattingParams represents the params body for the "SetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetFilenameFormatting.
*/
type SetFilenameFormattingParams struct {
	requests.Params

	// Filename formatting string to set.
	FilenameFormatting string `json:"filename-formatting"`
}

// GetRequestType returns the RequestType of SetFilenameFormattingParams
func (o *SetFilenameFormattingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetFilenameFormattingParams
func (o *SetFilenameFormattingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetFilenameFormattingParams
func (o *SetFilenameFormattingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetFilenameFormattingResponse represents the response body for the "SetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetFilenameFormatting.
*/
type SetFilenameFormattingResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetFilenameFormattingResponse
func (o *SetFilenameFormattingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetFilenameFormattingResponse
func (o *SetFilenameFormattingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetFilenameFormattingResponse
func (o *SetFilenameFormattingResponse) GetError() string {
	return o.Error
}

// SetFilenameFormatting sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetFilenameFormatting(
	params *SetFilenameFormattingParams,
) (*SetFilenameFormattingResponse, error) {
	params.RequestType = "SetFilenameFormatting"
	data := &SetFilenameFormattingResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
GetFilenameFormattingParams represents the params body for the "GetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetFilenameFormatting.
*/
type GetFilenameFormattingParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetFilenameFormattingParams
func (o *GetFilenameFormattingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetFilenameFormattingParams
func (o *GetFilenameFormattingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetFilenameFormattingParams
func (o *GetFilenameFormattingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetFilenameFormattingResponse represents the response body for the "GetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetFilenameFormatting.
*/
type GetFilenameFormattingResponse struct {
	requests.Response

	// Current filename formatting string.
	FilenameFormatting string `json:"filename-formatting"`
}

// GetMessageID returns the MessageID of GetFilenameFormattingResponse
func (o *GetFilenameFormattingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetFilenameFormattingResponse
func (o *GetFilenameFormattingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetFilenameFormattingResponse
func (o *GetFilenameFormattingResponse) GetError() string {
	return o.Error
}

// GetFilenameFormatting sends the corresponding request to the connected OBS WebSockets server.
// Note the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetFilenameFormatting(
	paramss ...*GetFilenameFormattingParams,
) (*GetFilenameFormattingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetFilenameFormattingParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetFilenameFormatting"
	data := &GetFilenameFormattingResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
